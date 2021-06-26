<template>
  <div>
    <table class="table">
      <NewsTableHeader />
      <tbody>
        <NewsItem
          v-for="(n, index) of news"
          v-bind:news="n"
          v-bind:index="index + 1"
          :key="n.id"
          v-on:remove-news="removeNews"/>
      </tbody>
    </table>
  </div>
</template>

<script>
import NewsTableHeader from "./NewsTableHeader";
import NewsItem from "./NewsItem";
export default {
  name: "NewsTable",
  props: ['news'],
  components: {
    NewsTableHeader,
    NewsItem
  },
  methods: {
    removeNews(id) {
      this.$emit('remove-news', id);
    }
  }
}
</script>

<style scoped>

</style>