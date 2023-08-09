// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  name: string
}

//post 接口
export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  // 输出 post 请求的 body
  console.log('request body ===>', req.body)
  res.status(200).json({ name: 'John Doe' })
}
