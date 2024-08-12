<template>
  <div class="flex flex-row justify-center m-10">
    <a-list
      class="list-demo-action-layout w-1/2"
      :bordered="false"
      :data="dataSource"
      max-height="500px"
      @reach-bottom="dataSource?.push(...dataSource)"
    >
      <template #item="{ item }">
        <a-list-item class="list-demo-item" action-layout="vertical" @click="clickVideo(item.bvid)">
          <a-list-item-meta
            :title="item.title"
          >
          </a-list-item-meta>
          <template #actions>
            <span @click="clickUp(item.mid)"><icon-heart />{{ item.up_name }}</span>
          </template>
          <template #extra>
            <div className="image-area">
              <img alt="pic" :src="item.pic" />
            </div>
          </template>
      </a-list-item>
      </template>
    </a-list>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { IconHeart } from '@arco-design/web-vue/es/icon';
import { tmpVideo, Vgroup, Video } from '@/types/apiType';
import { getVideo } from '@/utils/getData';

const dataSource = ref<Video[]>([]);

const clickVideo = (bvid: string) => {
  window.open(`https://www.bilibili.com/video/${bvid}`);
};

const clickUp = (mid: number) => {
  window.open(`https://space.bilibili.com/${mid}`);
};

onMounted(() => {
  const vg: Vgroup[] = localStorage.getItem('selectedVgroup') ? JSON.parse(localStorage.getItem('selectedVgroup')!) : [];
  console.log('%c [ vg ]-48', 'font-size:13px; background:#4337fe; color:#877bff;', vg);
  // TODO： type

  getVideo({list: vg}).then((res) => {
    dataSource.value = res;
  });
});
</script>

<style>
.list-demo-action-layout .image-area {
  width: 183px;
  height: 119px;
  border-radius: 2px;
  overflow: hidden;
  margin-left: 10px;
}

.list-demo-action-layout .list-demo-item {
  padding: 20px 0;
  border-bottom: 1px solid var(--color-fill-3);
  background-color: var(--color-fill-1);
  opacity: 0.7;
}

.list-demo-action-layout .list-demo-item:hover {
  background-color: var(--color-fill-2);
}

.list-demo-action-layout .list-demo-item:active {
  background-color: var(--color-fill-3);
}

.list-demo-action-layout .image-area img {
  width: 100%;
}

.list-demo-action-layout .arco-list-item-action .arco-icon {
  margin-right: 4px;
}

.arco-list-item-meta-title {
  font-size: 17px;
  font-weight: 500;
}


</style>
