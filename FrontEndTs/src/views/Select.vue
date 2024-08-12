// TODO: 设置组别并记忆选择
<template>
  <div class="w-full flex justify-center mt-10">
      <a-transfer 
        show-search
        one-way
        simple
        :data="data"
        v-model:model-value="selected"
        :title="['未选择', '已选择']"
        :source-input-search-props="{
          placeholder:'source search'
        }"
        :target-input-search-props="{
          placeholder:'target search'
        }"
      />
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, watch } from 'vue';
import { allVgroup } from '@/utils/getData';
import { tmpReq, Vgroup } from '@/types/apiType';

let vg: Vgroup[] = [];
const data = ref<{value: number, label: string}[]>([]);
const selected = ref(
  JSON.parse(localStorage.getItem('selectedVgroup') || '[]')
  .map(item => item.mid)
);

const getTransferData = async () => {
  vg = await allVgroup();
  // const vg = tmpReq.list;
  data.value = vg
  .sort((a, b) => b.label - a.label)
  .map((item) => {
    return {
      value: item.mid,
      label: item.uname || item.title
    };
  }) ?? [];
};

const selectedVgroup = computed(() => {
  console.log('%c [ vg ]-47', 'font-size:13px; background:#09a78a; color:#4debce;', selected.value);
  
  return selected.value.map((item) => {
    return vg.find((v) => v.mid === item);
  });
});

watch(selected, () => {
  localStorage.setItem('selectedVgroup', JSON.stringify(selectedVgroup.value));
})

onMounted(() => {
  getTransferData();
})
// computed + 异步 不行？
</script>

<style>
  div.arco-transfer-view {
    height: 500px;
    width: 250px;
  }
</style>