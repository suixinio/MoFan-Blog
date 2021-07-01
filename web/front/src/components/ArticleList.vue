<template>
  <v-col>
    <v-card class="ma-3" v-for="item in artList" :key="item.ID" link @click="$router.push(`detail/${item.ID}`)">
      <v-row no-gutters>
        <v-col class="d-flex justify-center align-center mx-3" cols="1">
          <v-img :src="item.img" max-width="100" max-height="100"></v-img>
        </v-col>
        <v-col>
          <v-card-title class="my-2">
            <v-chip color="pink" label class="white--text mr-2">{{ item.Category.name }}</v-chip>
            {{ item.title }}
          </v-card-title>
          <v-card-subtitle v-text="item.desc"></v-card-subtitle>
          <v-divider></v-divider>
          <v-card-text>
            <v-icon>{{ 'mdi-calendar-month' }}</v-icon>
            <span>{{ item.CreatedAt  | dateformat('YYY-MM-DD HH:SS') }}</span>
          </v-card-text>
        </v-col>
      </v-row>
    </v-card>
    <div class="text-center">
      <v-pagination
        v-model="queryParam.pagenum"
        :length="Math.ceil(this.total/queryParam.pagesize)"
        total-visible="7"
        @input="getArtList()"
      ></v-pagination>
    </div>
  </v-col>
</template>

<script>
export default {
  name: 'ArticleList',
  data () {
    return {
      artList: [],
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },
      total: 0
    }
  },
  created () {
    this.getArtList()
  },
  methods: {
    async getArtList () {
      const { data: res } = await this.$http.get('article', {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      this.artList = res.data
      this.total = res.total
      console.log(this.artList)
      console.log(this.total)
    }
  }
}
</script>

<style scoped>

</style>
