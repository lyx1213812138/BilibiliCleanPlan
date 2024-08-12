interface Vgroup {
  type: VgType;
  mid: number;
  uname?: string; // up name
  title?: string; // season title
  sid?: number;
  label: number; // 注意带上 label
}

enum VgType {
  IsSeason = 0,
  IsUp = 1
}

interface RequestMessage {
  list: Vgroup[];
}

interface Video {
  bvid: string;
  title: string;
  pic: string;
  mid: number;
  up_name: string;
}

const tmpReq: RequestMessage = {
  list: [
    {
      type: VgType.IsSeason,
      title: 'season title',
      mid: 439465191,
      sid: 1954425,
      label: 4
    },
    {
      type: VgType.IsUp,
      uname: 'up_name',
      mid: 8014168,
      label: 4
    }
  ]
};

const tmpVideo: Video = {
  bvid: 'BV19r4y1X7qN',
  title: '「Github一周热点35期」Docker中运行Mac，文生视频模型等5个项目',
  pic: 'http://i0.hdslb.com/bfs/archive/ec237020a592a6d2cf5fa9c53ebbe8e467fc5674.jpg',
  up_name: 'up_name',
  mid: 439465191
};

export { RequestMessage, Vgroup, VgType, Video, tmpReq, tmpVideo };