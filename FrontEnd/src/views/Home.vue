<template>
  <div class="flex flex-row justify-center m-10">
    <a-list
      class="list-demo-action-layout w-1/2"
      :bordered="false"
      :data="dataSource"
      max-height="500px"
      @reach-bottom="dataSource.push(...dataSource)"
    >
      <template #item="{ item }">
        <a-list-item class="list-demo-item" action-layout="vertical">
          <template #actions>
            <span><icon-heart />item.index</span>
            <span><icon-star />{{ item.index }}</span>
            <span><icon-message />Reply</span>
          </template>
          <template #extra>
            <div className="image-area">
              <img alt="arco-design" :src="item.imageSrc" />
            </div>
          </template>
          <a-list-item-meta
            :title="item.title"
            :description="item.description"
          >
          </a-list-item-meta>
        </a-list-item>
      </template>
    </a-list>
  </div>
  
</template>

<script setup>
import { ref } from 'vue';
import { IconHeart, IconStar, IconMessage } from '@arco-design/web-vue/es/icon';

const dataSource = ref(new Array(15).fill(null).map((_, index) => {
  return {
    index: index
  };
}));
</script>

<style scoped>
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

.list-demo-action-layout .image-area img {
  width: 100%;
}

.list-demo-action-layout .arco-list-item-action .arco-icon {
  margin: 0 4px;
}
</style>
