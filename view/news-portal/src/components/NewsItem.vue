<template>
  <tr>
    <td>{{ index }}</td>
    <td>{{ news.title }}</td>
    <td>{{ news.text }}</td>
    <td>{{ news.authors
        .map(a => `${a.firstName} ${a.lastName}`)
        .join(', ') }}
    </td>
    <td>
      <button type="button"
              class="btn btn-danger"
              @click="$emit('remove-news', news.id)">&times;</button>
    </td>
  </tr>
</template>

<script>
export default {
  props: ['news', 'index'],
  name: "NewsItem"
}
</script>

<style scoped>

</style>