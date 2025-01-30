export default {
    async fetch(request, env) {
        const { GITHUB_TOKEN, GITHUB_OWNER, } = env;
        const X_GITHUB_TOKEN = GITHUB_TOKEN;
        const X_GITHUB_OWNER = GITHUB_OWNER.split('/')[0]
        const X_GITHUB_REPO = GITHUB_OWNER.split('/')[1]

        if (!X_GITHUB_TOKEN || !X_GITHUB_OWNER || !X_GITHUB_REPO) {
            return new Response(undefined, { status: 424 });
        }

        if (request.method === 'POST') {
            const { headers } = request;

            const downloadUrl = headers.get('x-download-url');
            const dir = headers.get('x-dir') || '';
            const override = (headers.get('x-override') || 'false') === 'true';
            let fileName = headers.get('x-file-name');

            if (!downloadUrl) {
                return new Response('Missing x-download-url header', { status: 400 });
            }

            if (!fileName) {
                const fileId = `${downloadUrl}`.split('/').at(-1);
                fileName = `${dir}`.endsWith('/') ? `${dir}${fileId}` : `${dir}/${fileId}`;

                if (fileName.startsWith('/')) {
                    fileName = fileName.slice(1);
                }
            }

            try {
                const response = await fetch(downloadUrl);
                if (response.status != 200) {
                    throw new Error(`Failed to fetch: ${downloadUrl}, status code: ${response.status}`)
                }

                const fileContent = await response.arrayBuffer();
                const body = {
                    message: 'Uploaded file',
                    content: btoa(arrayBufferToString(fileContent)),
                }

                if (override) {
                    const getResponse = await getGitHubFile(X_GITHUB_OWNER, X_GITHUB_REPO, X_GITHUB_TOKEN, fileName);

                    if (getResponse.status === 200) {
                        body.sha = (await getResponse.json()).sha;
                    }
                }

                const githubUrl = `https://api.github.com/repos/${X_GITHUB_OWNER}/${X_GITHUB_REPO}/contents/${fileName}`;
                const githubResponse = await fetch(githubUrl, {
                    method: 'PUT',
                    headers: {
                        Authorization: `Bearer ${X_GITHUB_TOKEN}`,
                        'Content-Type': 'application/json',
                        'X-GitHub-Api-Version': ' 2022-11-28',
                        'User-Agent': 'friendly-foo-333',
                    },
                    body: JSON.stringify(body),
                });

                if (!githubResponse.ok) {
                    const errorData = await githubResponse.json();
                    throw new Error(`GitHub API error: ${errorData.message}`);
                }

                return new Response(undefined, {
                    headers: {
                        'location': fileName,
                    }, status: 302
                });
            } catch (error) {
                reportToTelegram(env, downloadUrl, error)
                return new Response(`Error uploading file: ${error}`, { status: 500 });
            }
        } else {
            return new Response('Method not allowed', { status: 405 });
        }
    },
};

function getGitHubFile(owner, repo, token, path) {
    const githubUrl = `https://api.github.com/repos/${owner}/${repo}/contents/${path}`;
    const githubResponse = fetch(githubUrl, {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${token}`,
            'Content-Type': 'application/json',
            'X-GitHub-Api-Version': ' 2022-11-28',
            'User-Agent': 'friendly-foo-333',
        },
    });

    return githubResponse;
}

function arrayBufferToString(buffer) {
    var binary = '';
    var bytes = new Uint8Array(buffer);
    var len = bytes.byteLength;
    for (var i = 0; i < len; i++) {
        binary += String.fromCharCode(bytes[i]);
    }
    return binary;
}

async function reportToTelegram(env, url, errorMessage) {
    const TELEGRAM_BOT_TOKEN = env.TELEGRAM_BOT_TOKEN;
    const TELEGRAM_CHAT_ID = env.TELEGRAM_CHAT_ID;
    const message = `Failed URL: ${url}\nError: ${errorMessage}`;

    const TELEGRAM_API = `https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage`;

    await fetch(TELEGRAM_API, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            chat_id: TELEGRAM_CHAT_ID,
            text: message
        })
    });
}