// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import { kv } from "@vercel/kv";


type Data = {
  name: string
}

//post 接口
export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  // 输出 post 请求的 body
  console.log('request body ===>', req.body)

  await kv.set("user_1_session", "session_token_value");
  const session = await kv.get("user_1_session");

  res.status(200).json({ name: 'John Doe' })
}
