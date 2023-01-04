# moodboard

![](https://github.com/felicepng/moodboard/blob/main/demo.gif)

A personal project to familiarise myself with goroutines, [SolidJS](https://www.solidjs.com), [Vite](https://vitejs.dev), as well as dabble with [OpenAI API](https://openai.com/blog/openai-api/).

## Implementation

1. The user submits a theme for the mood board via the user interface.
2. [OpenAI Text Completion API](https://beta.openai.com/docs/guides/completion) is used to generate a text response, consisting of a list of various image prompts referencing the given theme.
3. The list of prompts is parsed and [OpenAI Image Generation API](https://beta.openai.com/docs/guides/images) is used to create images for each prompt. This is done concurrently in the backend server using goroutines.
4. The images and prompts are returned to the frontend for displaying.

## Set-up

The OpenAI API uses API keys for authentication. Visit the [API Keys](https://beta.openai.com/account/api-keys) page to retrieve your API key and ensure that your account has sufficient credits for API usage before proceeding.

Create a `.env` file in the root directory and add your API key:

```bash
# .env
API_KEY=${YOUR_API_KEY}
```

Run the following commands to containerize the application with Docker:

```bash
docker-compose build
docker-compose up
```

Alternatively, you may run the `client` and `server` locally using the following commands:

```bash
# client
npm install
npm run dev
```

```bash
# server
go mod download
go run main.go
```

Open [http://localhost:3000](http://localhost:3000) to view the application in your browser.

## Testing

In the `server` directory, run `go test` to run all tests.
