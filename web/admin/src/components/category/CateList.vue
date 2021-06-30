<template>
  <div>
    <h3>分类列表页面</h3>
    <a-card>
      <a-row :gutter='20'>
        <a-col :span='6'>
          <a-input-search v-model='queryParam.name' placeholder='输入分类名查找' enter-button @search='getCateList'
                          allowClear />
        </a-col>
        <a-col :span='4'>
          <a-button type='primary' @click='addCateVisible = true'>新增分类</a-button>
        </a-col>
      </a-row>

      <a-table
        rowKey='ID'
        :columns='columns'
        :pagination='pagination'
        :dataSource='catelist'
        bordered
        @change='handleTableChange'
      >
        <template slot='action' slot-scope='data'>
          <div class='actionSlot'>
            <a-button type='primary' icon='edit' size='small' style='margin-right: 15px' @click='editCate(data.ID)'>编辑
            </a-button>
            <a-button type='danger' icon='delete' size='small' @click='deleteCate(data.ID)'>删除</a-button>
          </div>
        </template>
      </a-table>

    </a-card>
    <!--        新增分类区域 -->
    <a-modal
      closable
      title='新增用户'
      :visible='addCateVisible'
      width='60%'
      @ok='addCateOk'
      @cancel='addCateCancel'
      destroyOnClose
    >
      <a-form-model :model='newCate' :rules='addCateRules' ref='addCateRef'>

        <a-form-model-item label='分类名' prop='name'>
          <a-input v-model='newCate.name'></a-input>
        </a-form-model-item>

      </a-form-model>
    </a-modal>
    <!--    编辑分类-->
    <a-modal
      closable
      title='编辑分类'
      :visible='editCateVisible'
      width='60%'
      @ok='editCateOk'
      @cancel='editCateCancel'
      destroyOnClose
    >
      <a-form-model :model='cateInfo' :rules='cateRules' ref='editCateRef'>

        <a-form-model-item label='用户名' prop='name'>
          <a-input v-model='cateInfo.name'></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: '分类名',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center'
  },
  {
    title: '操作',
    width: '30%',
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
      columns,
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },
      addCateVisible: false,
      editCateVisible: false,
      cateInfo: {
        id: 0,
        name: ''
      },
      newCate: {
        name: ''
      },
      cateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.cateInfo.name === '') {
                callback(new Error('请输入分类名'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      addCateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.newCate.name === '') {
                callback(new Error('请输入分类名'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      }
    }
  },
  created () {
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
      this.getCateList()
    },
    async getCateList () {
      const { data: res } = await this.$http.get('category', {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        // eslint-disable-next-line no-constant-condition
        if (res.status === 1004 || res.status === 1005 || res.status === 1006 || res.status === 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.catelist = res.data
      this.pagination.total = res.total
    },
    // 删除分类
    deleteCate (id) {
      this.$confirm({
        title: '提示：请再次确认？',
        content: '确定要删除该分类吗？一旦删除，无法恢复',
        onOk: async () => {
          const res = await this.$http.delete(`category/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.getCateList()
          this.$message.success('删除成功')
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 新增分类
    addCateOk () {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('category/add', {
          name: this.newCate.name
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.addCateRef.resetFields()
        this.addCateVisible = false
        this.$message.success('添加分类成功')
        await this.getCateList()
      })
    },
    addCateCancel () {
      this.$refs.addCateRef.resetFields()
      this.addCateVisible = false
    },

    // 编辑用户
    editCateOk () {
      this.$refs.editCateRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`category/${this.cateInfo.id}`, {
          name: this.cateInfo.name
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.editCateVisible = false
        this.$message.success('更新分类成功')
        this.$refs.editCateRef.resetFields()
        this.getCateList()
      })
    },
    editCateCancel () {
      this.$refs.editCateRef.resetFields()
      this.editCateVisible = false
      this.$message.info('编辑已取消')
    },
    async editCate (id) {
      const { data: res } = await this.$http.get(`category/${id}`)
      this.editCateVisible = true
      this.cateInfo = res.data

      this.cateInfo.id = id
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
