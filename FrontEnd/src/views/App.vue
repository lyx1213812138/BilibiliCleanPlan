<template>
  <div
    :style="{
      backgroundImage: `url(${backgroundImageUrl})`,
      width: maxwidth + 'px',
      height: maxhight + 'px'
    }"
    class="bg-cover"
    id="container"
  > <!-- container of all, for background img-->
    <a-menu 
      mode="horizontal" 
      theme="dark" 
      :default-selected-keys="['home']" 
      class="opacity-80"
      id="menu"
    >
      <a-menu-item key="home">
        <router-link to="/">Home</router-link>
      </a-menu-item>
      <a-menu-item key="category">
        <router-link to="/category">Category</router-link>
      </a-menu-item>
    </a-menu>
    <router-view />
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
const maxwidth = ref(window.innerWidth);
const maxhight = ref(window.innerHeight);
const backgroundImageUrl = ref('/src/assets/img/background2.png');
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
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 1px;
}
</style>