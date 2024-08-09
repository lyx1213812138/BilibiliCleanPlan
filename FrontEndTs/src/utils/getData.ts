// TODO: 用store存储Vgroup数据
import axios from 'axios';

export function get(url: string, data: any) {
  // 用axios发送请求
  const str = JSON.stringify(data);
  axios.post(url, str).then((res) => {
    console.log(res.statusText);
  });
}

interface Vgroup {
  type: VgType;
  UpId: number;
  SeasonId?: number;
}

enum VgType {
  IsSeason = 0,
  IsUp = 1
}

interface RequestMessage {
  list: Vgroup[];
}

const tmp: RequestMessage = {
  list: [
    {
      type: VgType.IsSeason,
      UpId: 439465191,
      SeasonId: 1954425
    },
    {
      type: VgType.IsUp,
      UpId: 8014168
    }
  ]
};

get('http://localhost:12121', tmp);