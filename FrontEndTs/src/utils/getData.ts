// TODO: 用store存储Vgroup数据
import axios from 'axios';
import { RequestMessage, Vgroup, Video } from '@/types/apiType';

const url = 'http://localhost:12121';
const pathGetVideo = '/getvideo';
const pathAllVgroup = '/allvgroup';

export const getVideo = async (data: RequestMessage) => {
  const str = JSON.stringify(data);
  const resp = await axios.post(url + pathGetVideo, str);
  if (resp.status !== 200) {
    console.log('getVideo error', resp);
    return [];
  }
  if (typeof resp === 'object' && typeof resp.data === 'object') {
    const videos: Video[] = resp.data;
    return videos;
  } else {
    return [];
  }
}

export const allVgroup = async () => {
  const resp = await axios.get(url + pathAllVgroup);
  if (typeof resp === 'object' && typeof resp.data === 'object') {
    return resp.data as Vgroup[];
  } else {
    return [];
  }
}

export const handleRefresh = async () => {
  const resp = await axios.get(url + '/refresh');
  if (resp.status !== 200) {
    console.log('handleRefresh error', resp);
  }
}