<template>
  <div>
    <h3>文章列表页面</h3>
    <a-card>
      <a-row :gutter='20'>
        <a-col :span='6'>
          <a-input-search v-model='queryParam.title' placeholder='输入用户名查找' enter-button @search='getArtList'
                          allowClear />
        </a-col>
        <a-col :span='4'>
          <a-button type='primary' @click='$router.push(`/admin/addart/`)'>新增</a-button>
        </a-col>
        <a-col :span='3'>
          <a-select placeholder='请选择分类' style='width: 200px' @change='cateChange'>
            <a-select-option
              v-for='item in catelist'
              :key='item.ID'
              :value='item.ID'
            >{{ item.name }}
            </a-select-option>
          </a-select>
        </a-col>

      </a-row>

      <a-table
        rowKey='ID'
        :columns='columns'
        :pagination='pagination'
        :dataSource='artlist'
        bordered
        @change='handleTableChange'
      >
        <span slot='img' slot-scope='img' class='ArtImg'>
          <img :src='img'>
        </span>
        <template slot='action' slot-scope='data'>
          <div class='actionSlot'>
            <a-button type='primary' icon='edit' size='small' style='margin-right: 15px'
                      @click='$router.push(`/admin/addart/${data.ID}`)'>编辑
            </a-button>
            <a-button type='danger' icon='delete' size='small' @click='deleteArt(data.ID)'>删除</a-button>
          </div>
        </template>
      </a-table>

    </a-card>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '5%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类',
    dataIndex: 'Category.name',
    width: '5%',
    key: 'Category.name',
    align: 'center'
  },
  {
    title: '文章标题',
    dataIndex: 'title',
    width: '15%',
    key: 'title',
    align: 'center'
  },
  {
    title: '文章描述',
    dataIndex: 'desc',
    width: '20%',
    key: 'desc',
    align: 'center'
  },
  {
    title: '缩略图',
    dataIndex: 'img',
    width: '20%',
    key: 'img',
    align: 'center',
    scopedSlots: { customRender: 'img' }
  },
  {
    title: '操作',
    width: '15%',
    key: 'action',
    align: 'center',
    scopedSlots: { customRender: 'action' }
  }
]
export default {
  data () {
    return {
      pagination: {
        pageSizeOptions: ['5', '10', '20'],
        pageSize: 5,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `共${total}条`
      },
      catelist: [],
      artlist: [],
      columns,
      queryParam: {
        title: '',
        pagesize: 5,
        pagenum: 1
      }
    }
  },
  created () {
    this.getArtList()
    this.getCateList()
  },
  methods: {
    // 更改分页
    handleTableChange (pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getArtList()
    },
    // 获取分类列表
    async getCateList () {
      const { data: res } = await this.$http.get('category')
      if (res.status !== 200) {
        // eslint-disable-next-line no-constant-condition
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.catelist = res.data
      this.pagination.total = res.total
    },
    // 获取文章列表
    async getArtList () {
      const { data: res } = await this.$http.get('article', {
        params: {
          title: this.queryParam.title,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        // eslint-disable-next-line no-constant-condition
        if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
          window.sessionStorage.clear()
          await this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.artlist = res.data
      this.pagination.total = res.total
    },
    // 删除文章
    deleteArt (id) {
      this.$confirm({
        title: '提示：请再次确认？',
        content: '确定要删除该文章吗？一旦删除，无法恢复',
        onOk: async () => {
          const res = await this.$http.delete(`article/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.getArtList()
          this.$message.success('删除成功')
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 查询分类下的文章
    cateChange (value) {
      console.log(value)
      this.getCateAer(value)
    },
    async getCateAer (id) {
      const { data: res } = await this.$http.get(`article/list/${id}`)
      if (res.status !== 200) return this.$message.error(res.message)
      this.artlist = res.data
      this.pagination.total = res.total
    }

  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}

.ArtImg {
  height: 100%;
  width: 100%;
}

.ArtImg img {
  width: 100px;
  height: 80px;
}
</style>
