# demo bot with vercel

## Prerequisites

1. Create telegram bot using https://telegram.me/BotFather and save token
2. Set `TELEGRAM_TOKEN` env variable using command
```
export TELEGRAM_TOKEN=paste_token_from_step_1
```
3. Install go programming language https://golang.org/doc/install

# How to set webhook

```
curl https://api.telegram.org/bot{my_bot_token}/setWebhook?url={url_to_send_updates_to}
```
https://api.telegram.org/bot\2056915687:AAE0AMjFmd4ZY_C4mryMjJJf5ZfYyy1cLH0/setWebhook?url=https://demobot.vercel.app/api