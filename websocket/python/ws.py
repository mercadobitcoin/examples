#!/usr/bin/env python

import asyncio
from websockets import connect


async def process(msg):
    print(msg)

async def client(uri):
    async with connect(uri) as websocket:
        await websocket.send("{\"type\": \"subscribe\",\"subscription\": {\"name\": \"orderbook\",\"id\": \"BRLBTC\", \"limit\": 10}}")
        async for message in websocket:
            await process(message)


asyncio.run(client("wss://ws.mercadobitcoin.net/ws"))
