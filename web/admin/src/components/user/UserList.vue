<template>
  <div>
    <h3>用户列表页面</h3>
    <a-card>
      <a-row :gutter='20'>
        <a-col :span='6'>
          <a-input-search v-model='queryParam.username' placeholder='输入用户名查找' enter-button @search='getUserList'
                          allowClear />
        </a-col>
        <a-col :span='4'>
          <a-button type='primary' @click='addUserVisible = true'>新增</a-button>
        </a-col>
      </a-row>

      <a-table
        rowKey='ID'
        :columns='columns'
        :pagination='pagination'
        :dataSource='userlist'
        bordered
        @change='handleTableChange'
      >
        <span slot='role' slot-scope='role'>{{ role == 1 ? '管理员' : '订阅者' }}</span>
        <template slot='action' slot-scope='data'>
          <div class='actionSlot'>
            <a-button type='primary' icon='edit' size='small' style='margin-right: 15px' @click='editUser(data.ID)'>编辑
            </a-button>
            <a-button type='danger' icon='delete' size='small' @click='deleteUser(data.ID)'>删除</a-button>
          </div>
        </template>
      </a-table>

    </a-card>
    <!--        新增用户区域 -->
    <a-modal
      closable
      title='新增用户'
      :visible='addUserVisible'
      width='60%'
      @ok='addUserOk'
      @cancel='addUserCancel'
      destroyOnClose
    >
      <a-form-model :model='newUser' :rules='addUserRules' ref='addUserRef'>

        <a-form-model-item label='用户名' prop='username'>
          <a-input v-model='newUser.username'></a-input>
        </a-form-model-item>

        <a-form-model-item has-feedback label='密码' prop='password'>
          <a-input-password v-model='newUser.password'></a-input-password>
        </a-form-model-item>

        <a-form-model-item has-feedback label='确认密码' prop='checkpass'>
          <a-input-password v-model='newUser.checkpass'></a-input-password>
        </a-form-model-item>

        <a-form-model-item label='是否为管理员' prop='role'
        >
          <a-switch :checked='false' checked-children='是' un-checked-children='否' default-checked
                    @change='adminChange'></a-switch>

          <!--          <a-select defaultValue='2' style='width: 120px' @change='adminChange'>-->
          <!--            <a-select-option key='1' value.number='1'>是</a-select-option>-->
          <!--            <a-select-option key='2' value.number='2'>否</a-select-option>-->
          <!--          </a-select>-->
        </a-form-model-item>

      </a-form-model>
    </a-modal>
    <!--    编辑用户-->
    <a-modal
      closable
      title='编辑用户'
      :visible='editUserVisible'
      width='60%'
      @ok='editUserOk'
      @cancel='editUserCancel'
      destroyOnClose
    >
      <a-form-model :model='userInfo' :rules='userRules' ref='editUserRef'>

        <a-form-model-item label='用户名' prop='username'>
          <a-input v-model='userInfo.username'></a-input>
        </a-form-model-item>

        <a-form-model-item has-feedback label='密码' prop='password'>
          <a-input-password v-model='userInfo.password'></a-input-password>
        </a-form-model-item>

        <a-form-model-item has-feedback label='确认密码' prop='checkpass'>
          <a-input-password v-model='userInfo.checkpass'></a-input-password>
        </a-form-model-item>

        <a-form-model-item label='是否为管理员' prop='role'>
          <a-switch :checked='isAdmin' checked-children='是' un-checked-children='否' default-checked
                    @change='adminChange'></a-switch>
          <!--          <a-select defaultValue='2' style='width: 120px' @change='adminChange'>-->
          <!--            <a-select-option key='1' value.number='1'>是</a-select-option>-->
          <!--            <a-select-option key='2' value.number='2'>否</a-select-option>-->
          <!--          </a-select>-->
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
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    align: 'center',
    scopedSlots: { customRender: 'role' }
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
      userlist: [],
      columns,
      queryParam: {
        username: '',
        pagesize: 5,
        pagenum: 1
      },
      addUserVisible: false,
      editUserVisible: false,
      isAdmin: false,
      userInfo: {
        id: 0,
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      newUser: {
        username: '',
        password: '',
        role: 2,
        checkPass: ''
      },
      userRules: {
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.username === '') {
                callback(new Error('请输入用户名'))
              }
              if ([...this.userInfo.username].length < 4 || [...this.userInfo.username].length > 12) {
                callback(new Error('用户名应当在4到12个字符之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.password === '') {
                callback(new Error('请输入密码'))
              }
              if ([...this.userInfo.password].length < 6 || [...this.userInfo.password].length > 20) {
                callback(new Error('密码应当在6到20位之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.checkpass === '') {
                callback(new Error('请输入密码'))
              }
              if (this.userInfo.password !== this.userInfo.checkpass) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      addUserRules: {
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.username === '') {
                callback(new Error('请输入用户名'))
              }
              if ([...this.newUser.username].length < 4 || [...this.newUser.username].length > 12) {
                callback(new Error('用户名应当在4到12个字符之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.password === '') {
                callback(new Error('请输入密码'))
              }
              if ([...this.newUser.password].length < 6 || [...this.newUser.password].length > 20) {
                callback(new Error('密码应当在6到20位之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.checkpass === '') {
                callback(new Error('请输入密码'))
              }
              if (this.newUser.password !== this.newUser.checkpass) {
                callback(new Error('密码不一致，请重新输入'))
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
    this.getUserList()
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
      this.getUserList()
    },
    async getUserList () {
      const { data: res } = await this.$http.get('users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        // eslint-disable-next-line no-constant-condition
        if (res.status === 1004 || 1005 || 1006 || 1007) {
          window.sessionStorage.clear()
          this.$router.push('/login')
        }
        this.$message.error(res.message)
      }
      this.userlist = res.data
      this.pagination.total = res.total
    },
    // 删除用户
    deleteUser (id) {
      this.$confirm({
        title: '提示：请再次确认？',
        content: '确定要删除该用户吗？一旦删除，无法恢复',
        onOk: async () => {
          const res = await this.$http.delete(`user/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.getUserList()
          this.$message.success('删除成功')
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        }
      })
    },
    // 新增用户
    addUserOk () {
      console.log('role：' + this.userInfo.role)
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('user/add', {
          username: this.newUser.username,
          password: this.newUser.password,
          role: this.newUser.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.addUserVisible = false
        this.$message.success('添加用户成功')
        this.$refs.addUserRef.resetFields()
        await this.getUserList()
      })
    },
    addUserCancel () {
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
    },
    adminChange (checked) {
      if (checked) {
        this.userInfo.role = 1
      } else {
        this.userInfo.role = 2
      }
      // this.userInfo.role = value
    },
    // 编辑用户
    editUserOk () {
      console.log('role：' + this.userInfo.role)
      this.$refs.editUserRef.validate(async (valid) => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          password: this.userInfo.password,
          role: this.userInfo.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.editUserVisible = false
        this.$message.success('更新用户成功')
        this.$refs.editUserRef.resetFields()
        await this.getUserList()
      })
    },
    editUserCancel () {
      this.$refs.editUserRef.resetFields()
      this.editUserVisible = false
      this.isAdmin = false
      this.$message.info('编辑已取消')
    },
    async editUser (id) {
      const { data: res } = await this.$http.get(`user/${id}`)
      this.editUserVisible = true
      this.userInfo = res.data
      if (this.userInfo.role === 1) {
        this.isAdmin = true
      }
      this.userInfo.id = id
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
