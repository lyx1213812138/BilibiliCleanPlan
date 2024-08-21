<template>
  <!-- container of all, for background img-->
  <div
    :style="{
      backgroundImage: `url(${backgroundImageUrl})`,
      height: maxhight + 'px'
    }"
    class="bg-cover w-full"
    id="container"
  > 
    <!-- topbar -->
    <a-menu 
      mode="horizontal" 
      theme="dark" 
      :selected-keys="path"
      class="opacity-60"
      id="menu"
    >
      <a-menu-item key="/">
        <router-link to="/">Home</router-link>
      </a-menu-item>
      <a-menu-item key="/select">
        <router-link to="/select">Select</router-link>
      </a-menu-item>
    </a-menu>
    <router-view />

  </div>
</template>

<script setup lang="ts">
import { watch } from 'vue';
import { ref, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const maxwidth = ref(window.innerWidth);
const maxhight = ref(window.innerHeight);
const router = useRouter();
const path = ref([router.currentRoute.value.path]); // 加上[]失去了响应式
watch(() => router.currentRoute.value.path, (newPath) => {
  path.value = [newPath];
});

const backgroundImageUrl = ref('/background2.png');
</script>

<style>
#container {
  position: absolute;
  z-index: -2;
}

#container::before { /* 蒙版 */
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  background-color: rgba(0, 0, 0, 0.7); 
}

#menu::before { /* title */
  content: "Bilibili Clean Plan";
  position: absolute;
  width: 100%;
  height: 100%;
  font-size: 2.3em; 
  font-weight: 900;
  color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 1px;
}
</style>