<template>
    <main class="container">
        <div class="left-col">
            <div class="go-to-resume">
                <el-button
                    :icon="Promotion"
                    style="width: 200px;"
                    @click="goToTemplate" 
                    type="success" 
                    round 
                    size="large">去制作</el-button>
                    
            </div>
            <div class="ask-ai">
                <AiResume :btn_text = "'问问AI怎么写'"/>
            </div>
            <div class="section">
                <h4>刷题导航</h4>
                <div class="links">
                    <a href="https://leetcode.cn/" target="_blank">
                        <img src="../assets/leetcode.png" alt="LeetCode">
                        <span style="margin-left: -9px;">LeetCode</span>
                    </a>

                    <a href="https://www.nowcoder.com/" target="_blank">
                        <img src="../assets/cow.png" alt="cow">
                        <span>牛客网</span>
                    </a>
                </div>

                <div class="links">
                    <a href="https://www.mianshiya.com/" target="_blank">
                        <img src="../assets/duck.jpg" alt="duck" style="margin-left: -12px;">
                        <span>面试鸭</span>
                    </a>

                    <a href="https://www.luogu.com.cn/" target="_blank">
                        <img src="../assets/luogu.png" alt="luogu" style="margin-left: 16px;">
                        <span>洛谷</span>
                    </a>
                </div>

            </div>

            <div class="section">
                <h4>求职导航</h4>
                <div class="links">
                    <a href="https://www.zhipin.com/" target="_blank">
                        <img src="../assets/boss.jpg" alt="boss">
                        <span style="margin-left: -9px;">BOSS直聘</span>
                    </a>

                    <a href="https://www.shixiseng.com/" target="_blank">
                        <img src="../assets/shixi.png" alt="shixi">
                        <span>实习僧</span>
                    </a>
                </div>

                <div class="links">
                    <a href="https://www.ncss.cn/" target="_blank">
                        <img src="../assets/th.png" alt="th">
                        <span>国家大学生就业服务平台</span>
                    </a>

                </div>

            </div>

            <div class="section">
                <h4>大厂导航</h4>
                <div class="links">
                    <a href="https://www.huaweicloud.com/careers.html" target="_blank">
                        <img src="../assets/huawei.jpg" alt="huawei">
                        <span>华为云</span>
                    </a>

                    <a href="https://jobs.bytedance.com/" target="_blank">
                        <img src="../assets/byte.png" alt="zijie">
                        <span>字节跳动</span>
                    </a>
                </div>

                <div class="links">
                    <a href="https://careers.aliyun.com/campus/home?lang=zh" target="_blank">
                        <img src="../assets/ali.jpg" alt="ali" style="margin-left: -12px;">
                        <span>阿里云</span>
                    </a>

                    <a href="https://careers.tencent.com/home.html" target="_blank">
                        <img src="../assets/tencent.png" alt="tencent" style="margin-left: 5px;">
                        <span>腾讯云</span>
                    </a>

                </div>

            </div>
        </div>
        <div class="right-col">
            <h3 style="margin-left: 13vh;">求职经验荟萃</h3>
            <div class="exp-container">
                <div class="left-section">
                    <h4>灵感词汇</h4>
                    <div>
                        <el-select v-model="selectType" placeholder="请选择词汇类型" style="width: 80%">
                            <el-option label="动词" value="verb" />
                            <el-option label="形容词" value="adjective" />
                            <el-option label="量化词" value="weightWords" />
                        </el-select>
                    </div>
                    <div class="card-container">
                        <WordSection :words="selectedWords"/>
                    </div>
                </div>
                
                <div class="right-section">
                        <h4>经验汇总</h4>
                        <div class="demo-collapse">
                            <el-collapse v-model="activeName" accordion style="font-weight: bold;">
                            <el-collapse-item title="基本信息" name="1">
                                <div class="content">
                                姓名：确保姓名清晰且易于辨识，避免使用昵称。<br>
                                联系方式：提供有效的联系方式（手机号、邮箱）。确保邮箱地址专业，不要使用过于随意的邮箱地址。<br>
                                地址：地址可以根据需要选择是否包含，特别是在你申请的职位或公司有特定位置要求时。<br>
                                社交媒体链接（如LinkedIn、GitHub等）：如果相关，可以提供社交媒体链接，展示你的专业形象或作品。<br>
                                </div>
                            </el-collapse-item>
                            <el-collapse-item title="教育背景" name="2">
                                <div class="content">
                                    学历：列出最高学历并提供具体的学校名称、专业和毕业时间。<br>
                                    重要课程（可选）：列出与申请职位相关的重要课程或项目，尤其是对于应届毕业生或没有相关工作经验的人。<br>
                                    成绩和荣誉（可选）：如果有较高的GPA、奖学金或其他学术荣誉，可以加上。确保这些荣誉与职位相关。<br>
                                </div>
                            </el-collapse-item>
                            <el-collapse-item title="求职目的" name="3">
                                <div class="content">
                                    简短且针对性强。明确表达你的职业目标和你能为雇主带来的价值。<br>
                                    避免过于宽泛或不具体的陈述，例如“我想找一份我喜欢的工作”。相反，可以写成“希望利用我的项目管理和数据分析技能，在技术公司中推进战略发展。” <br>
                                </div>
                            </el-collapse-item>
                            <el-collapse-item title="项目经验" name="4">
                                <div class="content">
                                    如果你有相关的项目经验，尤其是和职位要求相关的，务必列出。<br>
                                    项目描述：简要介绍项目的背景、目标、你的角色以及所使用的技能。<br>
                                    量化成果：如果可能，突出项目的具体成果，例如“完成了一个移动应用程序，下载量超过1万次”。<br>
                                </div>
                            </el-collapse-item>
                            <el-collapse-item title="工作经历" name="5">
                                <div class="content">
                                    按照时间倒序排列工作经历，从最近的工作开始。<br>
                                    职位名称、公司名称和工作时间段必须清晰。<br>
                                    职责描述：<br>
                                    使用动词开头，如“管理”、“开发”、“推动”、“提升”等，突出你的贡献。<br>
                                    强调成果和影响，尽量用量化的指标来展示成果（例如“提升团队生产力30%”）。<br>
                                    尽量具体描述你在每个职位上做的工作，而不仅仅是列出职责。<br>
                                </div>
                            </el-collapse-item>
                            <el-collapse-item title="获奖荣誉" name="6">
                                <div class="content">
                                    列出你获得的与职位相关的证书或培训经历，例如编程语言认证、项目管理证书等。<br>
                                    提供证书的发证机构和获得日期，确保信息准确。<br>
                                </div>
                            </el-collapse-item>
                            <el-collapse-item title="技能证书" name="7">
                                <div class="content">
                                    硬技能：列出与职位相关的专业技能，如编程语言、数据分析工具、设计软件等。<br>
                                    软技能：包括沟通、领导力、团队合作等能力。<br>
                                    按照重要性和职位要求排序，并确保技能真实有效。<br>
                                </div>
                            </el-collapse-item>
                            </el-collapse>
                        </div>
                </div>
            </div>
        </div>
    </main>
</template>


<script setup>
import { useRouter } from 'vue-router'
import { Promotion } from '@element-plus/icons-vue'
import WordSection from './WordSection.vue';
import { ref, computed } from 'vue';
import AiResume from '../components/AiResume.vue';

const router = useRouter()

const goToTemplate = () => {
    router.push('/home/template')
  };

  const verbs = [
  '管理', '领导', '优化', '提升', '创新', '完成', '设计', '开发', '实现', '提供',
  '分析', '调整', '协作', '实施', '构建', '协调', '维护', '解决', '培训', '推动', '改进'
];

const adjectives = [
  '快速的', '创新的', '高效的', '灵活的', '强大的', '专业的', '出色的', '积极的', '高质量的', '精确的',
  '可靠的', '有效的', '完美的', '自信的', '优秀的'
];

const weightWords = ['增加','改善','减少','降低','增长','提升',
'优化','实现','达到','完成','节约','节省','百分比','缩短时间','缩短周期','提升效率','提高利润',
'提升销售额'] 

const selectType = ref('verb');

const selectedWords = computed(() => {
  if (selectType.value === 'verb') {
    return verbs;
  } else if (selectType.value === 'adjective') {
    return adjectives;
  } else if (selectType.value === 'weightWords') {
    return weightWords;
  }
  return []; // 默认返回空数组
});

</script>

<style scoped>
.content {
    white-space: pre-line;
    text-align: left;
}


.exp-container {
    display: flex;
    justify-content: space-between;
    padding: 20px;
    gap: 20px;
    height: 65vh;
}
.left-section  {
    width: 30%;
    text-align: center;
    margin-top: 10px;
    background-color: rgb(255, 255, 255);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    border: 0.8px solid rgba(255, 255, 255, 0.18);
    box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    -webkit-box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    border-radius: 12px;
    -webkit-border-radius: 12px;
    color: rgba(255, 255, 255, 0.75);
}

.right-section {
    width: 70%;
    text-align: center;
    margin-top: -7vh;
    background-color: rgb(255, 255, 255);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    border: 0.8px solid rgba(255, 255, 255, 0.18);
    box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    -webkit-box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    border-radius: 12px;
    -webkit-border-radius: 12px;
    color: rgba(255, 255, 255, 0.75);
    height: 75vh;
}

.demo-collapse {
    display: flex;
    justify-content: center;
}

.el-collapse {
    width: 80%;
}

h4 {
    color: black;
}

span {
    color: black;
}


.links {
    display: flex;
    justify-content: center; 
    gap: 10px; 
    margin-bottom: 15px; 
}

.links a {
    display: flex;
    text-decoration: none; /* 去掉链接的下划线 */
    color: black; /* 文字颜色 */
}

.links img {
    width: 30px; 
    height: 30px; 
    margin-right: 10px; 
}


.section {
    text-align: center;
    margin-top: 10px;
    background-color: rgb(255, 255, 255);
    backdrop-filter: blur(6px);
    -webkit-backdrop-filter: blur(6px);
    border: 0.8px solid rgba(255, 255, 255, 0.18);
    box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    -webkit-box-shadow: rgba(142, 142, 142, 0.19) 0px 6px 15px 0px;
    border-radius: 12px;
    -webkit-border-radius: 12px;
    color: rgba(255, 255, 255, 0.75);
}

.go-to-resume {
    margin-top: 20px;
    display: flex;
    justify-content: center;
}

.ask-ai {
    margin-top: 15px;
    display: flex;
    justify-content: center;
}

.left-col {
    box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
    width: 20%;
    float: left;
    height: 65vh;
    margin-left: 30px;
    
}

.right-col {
    box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0, 0, 0, 0.3) 0px 8px 16px -8px;
    width: 74%;
    float: right;
    height: 100vh;
    margin-right: 20px;
}

@media (min-width: 600px) {
    .left-col, .right-col {
        height: 85vh;
    }
}
</style>