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

  const newObj = {}

  // 遍历 post 请求的 body
  for (const key in req.body) {
    if (Object.prototype.hasOwnProperty.call(req.body, key)) {
      const element = req.body[key];
      // 先判断第二层是不是对象，
      // 将 post 请求的 body 第二层也循环遍历
      if (typeof element === 'object'){
        for (const key2 in element) {
          if (Object.prototype.hasOwnProperty.call(element, key2)) {
            const element2 = element[key2];
            // 将第二层的值是否为字符串
            // 然后判断长度是否大于 1000， 如果大于 1000 就截取前 1000 个字符

            if (typeof element2 === 'string') {
              if (element2.length > 100) {
                element[key2] = element2.slice(0, 100)
              } else {
                element[key2] = element2
              }
            }
          }
        }
      }

      newObj[key] = element
    }
  }
  console.log('request body ===>', newObj)

  // 保存 post 请求的 body 到 kv
  await kv.set("loging_data", JSON.stringify(newObj));

  res.status(200).json({ name: 'John Doe' })
}
