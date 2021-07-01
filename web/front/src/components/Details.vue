<template>
  <div>
    <div class="d-flex justify-center pa-3 text-h3 font-weight-bold">{{ artInfo.title }}</div>
    <div class="d-flex justify-center align-center pl-3 ml-3">
      <v-icon class="mr-1" small>{{ 'mdi-calendar-month' }}</v-icon>
      <span>{{ artInfo.CreatedAt  | dateformat('YYY-MM-DD HH:SS') }}</span>
    </div>

    <v-divider class="pa-3 ma-3"></v-divider>
    <v-alert class="ma-4" color="indigo" elevation="2" dark border="left" outlined>{{ artInfo.desc }}</v-alert>
    <div class="content text-justify ma-5 pa-3" v-html="artInfo.content"></div>
  </div>
</template>

<script>
export default {
  name: 'Details',
  props: ['id'],
  data () {
    return {
      artInfo: {}
    }
  },
  created () {
    this.getArtInfo()
  },
  methods: {
    // 查询文章
    async getArtInfo () {
      const { data: res } = await this.$http.get(`article/${this.id}`)
      if (res.status !== 200) {
        if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
          window.sessionStorage.clear()
          await this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.artInfo = res.data
      this.artInfo.id = res.data.ID
    }
  }
}
</script>

<style scoped>
.content >>> img, span, p {
  width: 100%;
}
</style>
