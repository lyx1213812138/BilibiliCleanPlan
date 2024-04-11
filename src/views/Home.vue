<template>
  <div class="flex flex-col content-center w-full"> <!-- 内容 -->
    <div class="w-2/3">aaa</div>
    <div class="w-2/3">bbb</div>
      <!-- <a-list
        :bordered="false"
        :data="contentData"
        class="w-2/3"
      >
        <template #item="{ item }">
          <a-list-item action-layout="vertical">
            <template #actions>
              <span><icon-heart />83</span>
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
              <template #avatar>
                <a-avatar shape="square">
                  <img alt="avatar" :src="item.avatar" />
                </a-avatar>
              </template>
            </a-list-item-meta>
          </a-list-item>
        </template>
      </a-list> -->
    </div>
</template>

<script setup>
import { ref } from "vue";
import { IconHeart, IconStar, IconMessage } from '@arco-design/web-vue/es/icon';

const contentData = ref([
  {
    index: 1,
    title: "Title 1",
    description: "Description 1",
  }
]);
</script>