// TODO: 用store存储Vgroup数据
import axios from 'axios';
import { RequestMessage, Video } from '@/types/apiType';

const url = 'http://localhost:12121';

export function getVideo(data: RequestMessage) {
  // 用axios发送请求
  const str = JSON.stringify(data);
  axios.post(url, str).then((res) => {
    const data: Video[] = JSON.parse(res.data);
    console.log("responce: ", data);
  });
}