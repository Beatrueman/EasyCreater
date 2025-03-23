<template>
    <main class="container">
        <div>
                <AlertMessage 
                v-model:visible="alertVisible"
                message="密码修改成功，请重新登录!"
                type="success"
                />
                
                <AlertMessage 
                v-model:visible="alertErrorVisible"
                :message="alertErrorMessage"
                type="error"
                />
            </div>
        <div class="left-col">
            <div>
                <el-button type="primary" class="back-button" @click="goBack"><el-icon><Back /></el-icon>返回</el-button>
            </div>
            <h3>个人信息</h3>
            <div class="info-container">
                <div class="section">
                    头像
                <AvatarUpload :sizeLimit="2"/>
                </div>
                <div class="section">
                    <el-icon><Avatar /></el-icon>
                    昵称：
                    {{ username }}
                </div>
                <div class="section">
                    <el-icon><PhoneFilled /></el-icon>
                    电话：
                    {{ phone }}
                </div>
                <div class="section">
                    <el-icon><Message /></el-icon>
                    邮箱：
                    {{ email }}
                </div>
            </div>
        </div>
        <div class="right-col">
            <h3 style="margin-top: 50px;">更改密码</h3>
            <div class="change-password-container">
                <div class="section"> 
                    <div>
                        <span style="color: black;">原密码</span>
                    </div>
                    <el-form-item prop="password">
                        <el-input
                            v-model="Password"
                            :prefix-icon="Lock"
                            show-password
                            style="width: 260px"
                            placeholder="请输入原密码" />
                    </el-form-item>
                </div>
                <div class="section"> 
                    <div>
                        <span style="color: black;">新密码</span>
                    </div>
                <el-form-item prop="password">
                        <el-input
                            v-model="newPassword"
                            :prefix-icon="Lock"
                            show-password
                            style="width: 260px"
                            placeholder="11-20位数字和字母组合" />
                    </el-form-item>
                </div>
                <el-button
                    style="margin-left: 170px;" 
                    type="success" 
                    class="back-button"
                    @click="changePassword">
                     <el-icon><Pointer /></el-icon>提交
                </el-button>
            </div>
        </div>
    </main>
</template>

<script setup>
  import { Lock } from '@element-plus/icons-vue'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getUserInfo, changeUserPassword } from '../apis/api'
import { ElMessage } from 'element-plus';
import AvatarUpload from '../components/AvatarUpload.vue';
import AlertMessage from '../components/AlertMessage.vue';

const router = useRouter()
const username = ref('')
const phone = ref('')
const email = ref('')

const Password = ref('')
const newPassword = ref('')

const alertVisible = ref(false);
const alertErrorVisible = ref(false);
const alertErrorMessage = ref('');


const changePassword = async () => {
    if (!Password.value || !newPassword.value) {
        alertErrorMessage.value = "请输入新密码";
        alertErrorVisible.value = true;
        setTimeout(() => { alertErrorVisible.value = false; }, 2000);
        return;
    }

    if (newPassword.value.length < 11 || newPassword.value.length > 20) {
        alertErrorMessage.value = "密码长度必须为11-20位";
        alertErrorVisible.value = true;
        setTimeout(() => { alertErrorVisible.value = false; }, 2000);
        return;
    }

    try {
        const response = await changeUserPassword({
            Password: Password.value,
            newPassword: newPassword.value
        });

        if (!response) {
            console.error("API 返回数据为空");
            throw new Error("API 返回数据为空"); 
        }

        if (response.status === 200) {
            alertVisible.value = true;
            setTimeout(() => {
                localStorage.removeItem('jwt-token');
                router.replace({ name: 'login' });
            }, 2000);
        } else {
            alertErrorMessage.value = response?.data?.msg || "修改密码失败!";
            alertErrorVisible.value = true;
            setTimeout(() => { alertErrorVisible.value = false; }, 2000);
        }
        } catch (error) {  
            console.error("请求失败:", error);

        if (error.response) {

            alertErrorMessage.value = error.response.data?.msg || "修改密码失败!";
        } else {
            alertErrorMessage.value = "服务器无响应，请稍后重试!";
        }

        alertErrorVisible.value = true;
        setTimeout(() => { alertErrorVisible.value = false; }, 2000);
}

};


onMounted(async () => {
    try {
        const response = await getUserInfo()

        if (!response) {
            throw new Error("API 返回数据为空");
        }

        if (response && response.status === 200) {
            username.value = response.username || "未知用户";
            phone.value = response.phone || "无";
            email.value = response.email || "无";
        } else {
            throw new Error(`用户信息获取失败: ${response.message || "未知错误"}`);
        }                           
    } catch (error) {
        console.error('获取用户信息失败:', error.message)
        alert('获取用户信息失败，请重新登录')
        router.replace({ name: 'login' })
    }
})

const goBack = () => {
    router.push('/home/index')
  }
</script>

<style scoped>

.alert-position {
    width: 500px;
}
.change-password-container {
    margin-top: 50px;
    background-color: rgb(255, 255, 255);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    border: 0.8px solid rgba(255, 255, 255, 0.18);
    box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    -webkit-box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    border-radius: 12px;
    -webkit-border-radius: 12px;
    color: rgba(255, 255, 255, 0.75);
    padding: 20px;
    width: 80%;
    margin-left: 30px;
}

.info-container {
    background-color: rgb(255, 255, 255);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    border: 0.8px solid rgba(255, 255, 255, 0.18);
    box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    -webkit-box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    border-radius: 12px;
    -webkit-border-radius: 12px;
    color: rgba(255, 255, 255, 0.75);
    padding: 20px;
    width: 80%;
    margin-left: 30px;
}

h3 {
    text-align: center;
}

.section {
    font-size: 18px;
    color: black;
    margin-left: 40px;
    margin-bottom: 30px;
    margin-top: 30px;
}
.left-col {
    box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
    width: 45%;
    float: left;
    height: 550px;
    margin-left: 30px;
}

.right-col {
    box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
    width: 45%;
    float: right;
    height: 550px;
    margin-right: 50px;
}

@media (min-width: 600px) {
    .left-col, .right-col {
        height: 80vh;
    }
}
</style>