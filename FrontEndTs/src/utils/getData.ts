// TODO: 用store存储Vgroup数据
import axios from 'axios';
import { RequestMessage, Video } from '@/types/apiType';

const url = 'http://localhost:12121';

export const getVideo = async (data: RequestMessage) => {
  const str = JSON.stringify(data);
  const videos: Video[] = JSON.parse((await axios.post(url, str)).data);
  return videos;
}