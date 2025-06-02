<template>
    <main class="container">
        <div v-if="is_display">
            <SideBar>
                <el-tabs tab-position="right" style="height: 100%;" class="demo-tabs">
                    <el-tab-pane label="编辑">
                    <ToggleSwitch 
                        :toggle-active="editing"
                        @switch-toggled="toggleEditMode" 
                        label="编辑模式" 
                        :defaultValue="false"/>   

                    <div>
                        <ExportPdf v-if="!editing" :resume-format="resumeFormat"/>
                    </div>

                    <div>
                        <span style="color: black;">简历格式</span>
                        <SelectInput 
                        :options="[{'name':'a4', 'value':'a4'}, {'name':'信件', 'value':'letter'}]"
                        :default-option="resumeFormat"
                        @update-selection="resumeFormat = $event"
                        />
                    </div>

                    <div v-if="editing">
                        <div style="color: black; font-weight: bolder">左列</div>
                        <ColorInput label="标题颜色" :default-color="colors.left.highlight" @color-changed="colors.left.highlight = $event" />
                        <ColorInput label="背景颜色" :default-color="colors.left.background" @color-changed="colors.left.background = $event" />
                        <ColorInput label="文本颜色" :default-color="colors.left.text" @color-changed="colors.left.text = $event" />
                        <PercentageInput label="左列宽度控制" :min="20" :max="80" :current-value="widthLeft" @percentage-changed="widthLeft = $event" />

                        <div style="color: black; font-weight: bolder">右列</div>
                        <ColorInput label="标题颜色" :default-color="colors.right.highlight" @color-changed="colors.right.highlight = $event" />
                        <ColorInput label="背景颜色" :default-color="colors.right.background" @color-changed="colors.right.background = $event" />
                        <ColorInput label="文本颜色" :default-color="colors.right.text" @color-changed="colors.right.text = $event" />
                        <span>标题粗细</span>
                        <SelectInput @update-selection="headlineWeight = $event" :default-option="headlineWeight" :options="[{'name': '细', 'value': '300'}, {'name': '正常', 'value': '400'}, {'name': '粗', 'value': '600'}]" />

                        <div style="color: black; font-weight: bolder">头像控制</div>
                        <ToggleSwitch @switch-toggled="toggleImageDisplay" label="显示头像" :defaultValue="true"/>
                        <div v-if="showImage">
                        <span style="color: black;">头像形状</span>
                        <SelectInput @update-selection="imageShape = $event" :default-option="imageShape" :options="[{'name': '方形', 'value': 'square'}, {'name': '圆形', 'value': 'circle'}]" />
                        <div>
                        <span style="color: black;">上传头像</span>
                        </div>
                        <ImageUpload @image-changed="imgUrl = $event" />
                        </div>
                    </div>
                    </el-tab-pane>

                    <el-tab-pane label="灵感">
                        <div class="infinite-scroll-wrapper" style="height: 500px; overflow: auto">
                            
                            <!-- 岗位相关 byRole -->
                            <h3 style="color: black;">岗位相关</h3>
                            <div v-for="(role, index) in words.byRole" :key="'role-' + index" class="mb-4">
                            <h4 style="color: black;">{{ role.job }}</h4>
                            <ul class="infinite-list">
                                <li v-for="(item, idx) in role.items" :key="'role-item-' + index + '-' + idx" class="infinite-list-item">
                                {{ item }}
                                </li>
                            </ul>
                            </div>

                            <!-- 动词 -->
                            <h3 style="color: black;">动词</h3>
                            <ul class="infinite-list">
                            <li style="color: black;" v-for="(word, index) in words.verbs" :key="'verb-' + index" class="infinite-list-item">
                                {{ word }}
                            </li>
                            </ul>

                            <!-- 名词 -->
                            <h3 style="color: black;">名词</h3>
                            <ul class="infinite-list">
                            <li style="color: black;" v-for="(word, index) in words.nouns" :key="'noun-' + index" class="infinite-list-item">
                                {{ word }}
                            </li>
                            </ul>

                            <!-- 形容词 -->
                            <h3 style="color: black;">形容词</h3>
                            <ul class="infinite-list">
                            <li style="color: black;" v-for="(word, index) in words.adjectives" :key="'adj-' + index" class="infinite-list-item">
                                {{ word }}
                            </li>
                            </ul>

                            <!-- 权重词 -->
                            <h3 style="color: black;">权重词</h3>
                            <ul class="infinite-list">
                            <li style="color: black;" v-for="(word, index) in words.weightWords" :key="'weight-' + index" class="infinite-list-item">
                                {{ word }}
                            </li>
                            </ul>

                        </div>
                    </el-tab-pane>


                </el-tabs>
            </SideBar>
            <div style="color: black">
                <el-button type="primary" style="margin-bottom: 20px; margin-top: 0px;" @click="goToTemplate"><el-icon><Back /></el-icon>返回</el-button>
                <CustomButton
                            v-if="!editing" 
                            btn-type="primary" 
                            style="margin-bottom: 20px; margin-top: 0px;"
                           @click="openDialog"
                            >保存简历</CustomButton>
                <AiPolish :fromTemplate="fromTemplate" @saveResumeLocal="saveResumeData"/>
                <el-dialog
                    v-model="dialogVisible"
                    title="请输入简历名称"
                    width="500"
                    :before-close="handleClose"
                >
                <el-input 
                    v-model="resumeName" 
                    placeholder="建议以岗位名称命名，如“算法工程师”" 
                    clearable 
                />
                    <template #footer>
                    <div class="dialog-footer">
                        <el-button @click="dialogVisible = false">取消</el-button>
                        <el-button type="primary" @click="confirmSave">
                        确定
                        </el-button>
                    </div>
                    </template>
                </el-dialog>

            </div>
        </div>
            <div 
                id="resume"
                class="d-flex" 
                :class="{ 'edit-off': !editing, 'letter-format': resumeFormat == 'letter' }"
                :style="cssVariables"
                >
                <div class="left-col">
                    <div 
                        class="name"
                        :contenteditable="editing" 
                        @blur="updateProperty($event, 'Myname')">
                        {{ Myname }}
                    </div>
                    <div 
                        class="job-title"
                        :contenteditable="editing"
                        @blur="updateProperty($event, 'title')">
                        {{ title }}</div>
                    <div class="resume-section">
                        <div class="d-flex">
                            <SectionHeadline :editing="editing" :headline="headlines[4]" @headline-edited="updateHeadline($event, 4)"/>
                            <EditButtons :editing="editing" :show-remove-button="false" @add-click="addEducation"/>
                        </div>
                            <div v-for="(item, index) in education" :key="index" class="inner-section relative-container">
                            <div class="d-flex justify-content-between">
                                <div :contenteditable="editing" @blur="updateEducation($event, 'title', index)">
                                {{ item.title }}
                            </div>
                            <div class="delete-button-container">
                                <EditButtons :editing="editing" @remove-click="removeEducation(index)" :show-add-button="false"/>
                            </div>
                            </div>
                            <div class="d-flex justify-content-between" style="font-weight: bolder;">
                                <div>
                                    <sapn :contenteditable="editing" @blur="updateEducation($event, 'university', index)">{{ item.university }}</sapn>
                                </div>
                                <div>
                                    <span :contenteditable="editing" @blur="updateEducation($event, 'major', index)">{{ item.major }}</span>
                                </div>
                                <div>
                                    <span :contenteditable="editing" @blur="updateEducation($event, 'date', index)">{{ item.date }}</span>
                                </div>
                            </div>
                            <div>主修课程：<span :contenteditable="editing" @blur="updateEducation($event, 'description', index)">{{ item.description }}</span></div>
                        </div>
                    </div>
                    <div class="resume-section">
                    <div class="d-flex">
                        <SectionHeadline :editing="editing" :headline="headlines[5]" @headline-edited="updateHeadline($event, 5)"/>
                        <EditButtons :editing="editing" :show-remove-button="false" @add-click="addProjectExperience"/>
                    </div>
                        <div v-for="(item, index) in project_experience" :key="index" class="inner-section relative-container">
                        <div class="d-flex justify-content-between" style="font-weight: bolder;">
                            <div :contenteditable="editing" @blur="updateProject($event, 'name', index)">{{ item.name }}</div>
                            <div class="delete-button-container">
                            <EditButtons :editing="editing" @remove-click="removeProjectExperience(index)" :show-add-button="false"/>
                        </div>
                        </div>
                            <div class="d-flex justify-content-between">
                            <div>技术栈：<span :contenteditable="editing" @blur="updateProject($event, 'tech_stack', index)">{{ item.tech_stack }}</span></div>
                        </div>
                        <ul>
                            <li v-for="(desc, innerIndex) in item.description" 
                            :key="innerIndex"
                            :contenteditable="editing"
                            @blur="updateProjectDescription($event, index, innerIndex)"
                            >{{ desc }}</li>
                        </ul>
                        <EditButtons 
                        :editing="editing"
                        @add-click="item.description.push('编辑项')" 
                        @remove-click="item.description.pop()"
                        :show-remove-button="item.description.length > 0"
                    />
                    </div>
                </div>
                    <div class="resume-section">
                        <div class="d-flex">
                            <SectionHeadline :editing="editing" :headline="headlines[6]" @headline-edited="updateHeadline($event, 6)"/>
                            <EditButtons :editing="editing" :show-remove-button="false" @add-click="addExperience"/>
                        </div>
                            <div v-for="(item, index) in experience" :key="index" class="inner-section relative-container">
                                <div class="d-flex justify-content-between" style="font-weight: bold;">
                                    <div :contenteditable="editing" @blur="updateExperience($event, 'position', index)">{{ item.position }}</div>
                                        <div class="delete-button-container">
                                            <EditButtons :editing="editing" @remove-click="removeExperience(index)" :show-add-button="false"/>
                                        </div>
                                    </div>
                                <div class="d-flex justify-content-between" style="font-weight: bold;">
                                <div>
                                    <span :contenteditable="editing" @blur="updateExperience($event, 'company', index)">{{ item.company }}</span>
                                </div>
                                <div>
                                    <span :contenteditable="editing" @blur="updateExperience($event, 'location', index)">{{ item.location }}</span>
                                </div>
                                <div>
                                    <span :contenteditable="editing" @blur="updateExperience($event, 'date', index)">{{ item.date }}</span>
                                </div>
                            </div>
                            <ul>
                                <li v-for="(desc, innerIndex) in item.description" 
                                :key="innerIndex"
                                :contenteditable="editing"
                                @blur="updateExperienceDescription($event, index, innerIndex)"
                                >
                                {{ desc }}</li>
                            </ul>
                            <EditButtons
                            :editing="editing" 
                            @add-click="item.description.push('编辑项')" 
                            @remove-click="item.description.pop()"
                            :show-remove-button="item.description.length > 0"
                        />
                        </div>
                    </div>
                </div>
                <div class="right-col" :style="{width: percentageWidthRight}">                    
                    <div class="resume-section">
                        <img :src="imgUrl"
                            v-if="showImage"
                            class="profile-pic" 
                            alt="profile picture"
                            :class="{'circle': imageShape == 'circle'}"
                            >
                        
                        <SectionHeadline :editing="editing" :headline="headlines[0]" @headline-edited="updateHeadline($event, 0)"/>
                        
                        <div 
                            :contenteditable="editing" 
                            @blur="updateProperty($event, 'introText')">
                            {{ introText }}
                        </div>
                    </div>
                    <div class="resume-section">
                        <SectionHeadline :editing="editing" :headline="headlines[1]" @headline-edited="updateHeadline($event, 1)"/>
                        
                        <Contact :editing="editing" :contact="contact" @edit="updateNestedProperty"/>
                        
                    </div>
                    <div class="resume-section">
                        <SectionHeadline :editing="editing" :headline="headlines[2]" @headline-edited="updateHeadline($event, 2)"/>
                        <ul>
                            <li v-for="(skill, index) in skills" :key="index" :contenteditable="editing" @blur="updateNestedProperty($event, 'skills', index)">{{ skill }}</li>
                        </ul>
                        <EditButtons
                            :editing="editing"  
                            @add-click="skills.push('编辑项')" 
                            @remove-click="skills.pop()"
                            :show-remove-button="skills.length > 0"
                        />
                    </div>
                    <div class="resume-section">
                        <SectionHeadline :editing="editing" :headline="headlines[3]" @headline-edited="updateHeadline($event, 3)"/>
                        <ul>
                            <li v-for="(h, index) in honor" :key="index" :contenteditable="editing" @blur="updateNestedProperty($event, 'honor', index)">{{ h }}</li>
                        </ul>
                        <EditButtons 
                            :editing="editing"
                            @add-click="honor.push('编辑项')" 
                            @remove-click="honor.pop()"
                            :show-remove-button="honor.length > 0"
                        />
                    </div>
                </div>
            </div>
    </main>
</template>

<style scoped>
.infinite-list {
  padding: 0;
  margin: 0;
  list-style: none;
}

.infinite-list .infinite-list-item {
  display: block;         /* 让每个 item 占一行 */
  padding: 10px 15px;
  margin: 5px 0;
  background-color: #f9f9f9;
  border-radius: 5px;
  color: #333;            /* 更好的可读性 */
  font-size: 14px;
  line-height: 1.5;
  text-align: left;       /* 左对齐更适合长句展示 */
}

.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}

.demo-tabs > .el-tabs__content {
  padding: 12px;
  overflow-y: auto;
}

.el-tabs--right .el-tabs__content {
  height: 100%;
  width: 20%;
}

@media (min-width: 1350px) {
    .resume {
        margin-right: 300px;
    }
}

@media (min-width: 1600px) {
    .resume {
        margin-left: auto;
        margin-right: auto;
    }
}

  :root {
    --highlight-color-left: #82C0CC;
    --background-color-left: #e0e0e0;
    --text-color-left: white;
    --highlight-color-right: #82C0CC;
    --background-color-right: white;
    --text-color-right: black;
    --headline-weight: 600;
  }

body {
    font-family: sans-serif;
    font-size: 16px;
}

.container {
    margin: 40px auto;
    max-width: 1200px;
}

.right-col {
    background-color: var(--background-color-left);
    color: var(--text-color-left);
    border-right: 1px sold var(--highlight-color-left);
    padding: 30px;
}

.left-col {
    background: var(--background-color-right);
    color: var(--text-color-right);
    width: 70%;
    padding: 30px;
}

.section-headline {
    font-size: 20px;
    font-weight: var(--headline-weight);
    margin-bottom: 15px;
    margin-top: 0;
}

.right-col .section-headline {
    color: var(--highlight-color-right);
}

.left-col .section-headline {
    color: var(--highlight-color-left);
    border-bottom: 1px solid var(--highlight-color-left);
    padding-bottom: 5px;
    margin-right: -30px;
    padding-right: 10px;
}

.resume-section {
    margin-bottom: 30px;
    line-height: 1.5;
}

.name {
    font-size: 28px;
    color: var(--highlight-color-left);
    border-bottom: 1px solid var(--highlight-color-left);
    margin: 0;
    margin-left: -30px;
    padding-left: 30px;
    padding-bottom: 15px;
}

.job-title {
    border-bottom: 1px solid var(--highlight-color-right);
    margin: 0 0 20px -30px;
    padding: 15px 0 15px 30px;
    font-weight: 300;
    font-size: 20px;
}

.d-flex {
    display: flex;
}
  #resume {
    box-shadow: rgba(50, 50, 93, 0.25) 0px 13px 27px 05px, rgba(0,0,0,0.3) 0px 8px 16px -8px;
    height: 297mm;
    width: 210mm;
    min-width: 800px;
  }

  #resume ul {
    padding-inline-start: 16px;
    margin-block-end: 5px;
    margin-block-start: 5px;
  }

.profile-pic {
    display: block;
    width: 160px;
    height: 160px;
    border: 5px solid var(--highlight-color-right);
    margin-bottom: 20px;
    object-fit: cover;
    margin-left: auto;
    margin-right: auto;
}

.circle {
    border-radius: 50%;
}

.inner-section {
    margin-bottom: 20px;
}

.justify-content-between {
    justify-content: space-between;
}

.relative-container {
    position: relative; 
    padding-right: 0px; 
}

.delete-button-container {
    position: absolute; 
    right: 0; 
    top: 0; 
}

[contenteditable="true"] {
    background-color: rgba(216, 216, 216, 0.253);
    padding: 2px;
    border-radius: 2px;
}

#resume.edit-off {
    height: 296.7mm;
    margin-left: -7.5px;
    margin-top: 0;
}

#resume.edit-off.letter-format {
    width: 8.5in;
    height: 11in;
}
</style>

<script>
import SectionHeadline from '../../components/SectionHeadline.vue';
import Contact from '../../components/Contact.vue';
import EditButtons from '../../components/EditButtons.vue';
import ToggleSwitch from '../../components/ToggleSwitch.vue';
import { fa } from 'element-plus/es/locales.mjs';
import SideBar from '../../components/SideBar.vue';
import ColorInput from '../../components/ColorInput.vue';
import PercentageInput from '../../components/PercentageInput.vue';
import SelectInput from '../../components/SelectInput.vue';
import { Select, Upload } from '@element-plus/icons-vue';
import ImageUpload from '../../components/ImageUpload.vue';
import ExportPdf from '../../components/ExportPdf.vue';
import CustomButton from '../../components/CustomButton.vue';
import { useRouter } from 'vue-router';
import AiPolish from '../../components/AiPolish.vue';
import { fetchResume, saveResume, getIdeas } from '../../apis/api';
import html2canvas from 'html2canvas';

const router = useRouter();

export default {
    created() {
        const resumeId = Number(this.$route.query.resume_id);
        
        if (resumeId === -1) {
            const currentTemplate = this.$route.query.template || 'template2'; // 从 query 中获取模板名
            this.templateName = currentTemplate;
            // 从本地加载简历
            const resumeData = localStorage.getItem(`resumeData_${currentTemplate}`);
            // 检查 templateName是否为template2
            if (resumeData) {
                const parsedData = JSON.parse(resumeData);
                this.templateName = parsedData.fromTemplate;
                if (this.templateName === 'template2') {
                    this.loadIntoData(parsedData);
                    console.log('从本地加载简历',this.templateName);
                } else {
                    console.log('本地简历数据不匹配，使用默认模板');
                }
            } else {
                console.log('本地没有简历数据');
            }
        } else {
            console.log("从数据库加载简历",resumeId);
            this.loadFromDatabase(resumeId);
        }

        const is_display_str = this.$route.query.is_display;
        if (is_display_str !== undefined) {
        // 将 is_display 转换为布尔值
        this.is_display = is_display_str === 'true'; // 'true' 字符串转换为 true
        }
    },
    components: {
        SectionHeadline,
        Contact,
        EditButtons,
        ToggleSwitch,
        SideBar,
        ColorInput,
        PercentageInput,
        SelectInput,
        ImageUpload,
        ExportPdf,
        CustomButton,
        AiPolish
    },
    data() {
        return {
            colors: {
                left: {
                    highlight: '#82C0CC',
                    text: 'black',
                    background: '#F9DF7A'
                },
                right: {
                    highlight: '#303030',
                    text: 'black',
                    background: 'white'
                },
            },
            Myname: "Yiiong",
            title: "SRE工程师",
            introText: "个人介绍",
            imgUrl: '/avatar.png',
            headlines: ["个人信息", "联系方式", "专业技能", "荣誉/证书", "教育经历", "项目经历", "工作经历"], // 修改了这里
            honor: ["国家一等奖学金","HCIE认证"],
            contact: {
                phone: "19823838711",
                email: "yiiong@qq.com",
                address: "重庆"
            },
            skills: [
                "Python",
                "Linux",
                "Git",
                "Docker",
                "Kubernetes",
                "Go",
            ],
            education: [
                {
                    title: "本科",
                    university: "重庆邮电大学",
                    major: "计算机科学与技术",
                    date: "2022-2026",
                    description: "主修课程：计算机网络、数据结构、操作系统、计算机组成原理"
                },
                {
                    title: "硕士",
                    university: "重庆邮电大学",
                    major: "计算机科学与技术",
                    date: "2022-2026",
                    description: "从事的研究方向"
                },
            ],
            experience: [{
                company: "ABC公司",
                position: "算法工程师",
                location: "重庆",
                date: "2022-至今",
                description: [
                    "1. 搭建云服务器，使用Python进行开发",
                    "2. 使用Python进行机器学习",
                    "3. 使用Python进行机器学习"
                ]
            }, {
                company: "XXX公司",
                position: "云计算工程师",
                location: "重庆",
                date: "2022-至今",
                description: [
                    "1. 搭建云服务器，使用Python进行开发",
                    "2. Kubernetes 部署"
                    ]
                },
            ],
            project_experience: [{
                name: "在线简历生成器",
                tech_stack: "Python、机器学习、Kubernetes、Docker、Linux、Git",
                description: [
                    "1. 使用Vue 开发",
                    "2. 后端使用Go 开发",
                    "3. 在 Kubernetes 部署"
                ]
            }],
            editing: false,
            showImage: true,
            widthRight: 30,
            imageShape: "circle",
            headlineWeight: "400",
            resumeFormat: 'a4',
            templateName: "template2",
            fromTemplate: "template2",
            aiResponse: "",    // AI 返回的数据
            loading: false,
            resumeId: null,
            is_display: true, 
            dialogVisible: false,
            resumeName: "",   
            count: 0,
            maxCount: 200,
            words: [],
        }
    },
    computed: {
        cssVariables() {
            return {
                '--highlight-color-left': this.colors.left.highlight,
                '--highlight-color-right': this.colors.right.highlight,
                '--text-color-left': this.colors.left.text,
                '--text-color-right': this.colors.right.text,
                '--background-color-left': this.colors.left.background,
                '--background-color-right': this.colors.right.background,
                '--headline-weight': this.headlineWeight,
            }
        },
        percentageWidthRight() {
            return this.widthRight + '%';
        },
        disabled() {
            return this.count >= this.maxCount
        }, 
    },
    methods: {
        load() {
        if (this.count < this.maxCount) {
            this.count += 2
            }
        },
        updateHeadline(newValue, index) {
            this.headlines[index] = newValue;
        },
        updateProperty(event, key) {
            this[key] = event.target.innerText;
        },
        updateNestedProperty(event, key1, key2) {
            this[key1][key2] = event.target.innerText;
        },
        updateExperience(event, key, index) {
            this.experience[index][key] = event.target.innerText
        },
        updateExperienceDescription(event, index1, index2) {
            this.experience[index1]['description'][index2] = event.target.innerText
        },
        updateEducation(event, key, index) {
            this.education[index][key] = event.target.innerText
        },
        updateEducationDescription(event, index1, index2) {
            this.experience[index1]['description'][index2] = event.target.innerText
        },
        updateProject(event, key, index) {
            this.project_experience[index][key] = event.target.innerText
        },
        updateProjectDescription(event, index1, index2) {
            this.project_experience[index1]['description'][index2] = event.target.innerText
        },
        addExperience() {
            this.experience.unshift({
                company: "xxx公司",
                position: "算法工程师",
                location: "重庆",
                date: "任职时间",
                description: ["工作描述"]
            })
        },
        addEducation() {
            this.education.unshift({
                title: "年级",
                university: "学校",
                major: "专业",
                date: "入学时间",
                description: "研究方向以及主修课程"
            })
        },
        addProjectExperience() {
            this.project_experience.unshift({
                name: "项目名称",
                tech_stack: "使用的技术栈",
                description: ["项目描述"]
            })
        },
        removeExperience(index) {
            this.experience.splice(index, 1)
        },
        removeEducation(index) {
            this.education.splice(index, 1)
        },
        removeProjectExperience(index) {
            this.project_experience.splice(index, 1)
        },
        toggleEditMode(value) {
            this.editing = value;
        },
        toggleImageDisplay(value) {
            this.showImage = value;
        },
        async saveConfig() {
            const { is_display, dialogVisible, ...resumeData } = this.$data; 

            try {
                // 生成缩略图
                const element = document.getElementById('resume');
                if (!element) {
                    console.error('未找到元素');
                    return;
                }

                const canvas = await html2canvas(element, { scale: 5, width: element.scrollWidth });
                const thumbnailDataUrl = canvas.toDataURL('image/jpeg'); // 获取base64

                // 添加缩略图
                await saveResume(resumeData, thumbnailDataUrl);
                localStorage.setItem('resumeData', JSON.stringify(resumeData));
                console.log('保存成功');
            } catch (error) {
                console.error('保存失败:', error);
            }
        },

        loadIntoData(config) {
            for(const key in config) {
                if(this.$data.hasOwnProperty(key) && key !== 'fromTemplate') {
                    this[key] = config[key]
                }
            }
        },

        async loadFromDatabase(resumeId) {
            console.log(`正在从数据库加载简历 ID: ${resumeId}`);
            try {
                const resume = await fetchResume(resumeId);
                if(resume) {
                    const parsedData = JSON.parse(resume[0].resume_data);
                    this.loadIntoData(parsedData);
                    
                    console.log('加载成功');
                } else {
                    console.log('简历不存在');
                }
            } catch (error) {
                console.error(`加载简历 ID ${resumeId} 失败:`, error);
            }
        },

        goToTemplate() {
            this.$router.push('/home/template');
        },

        openDialog() {
            this.resumeName = "";
            this.dialogVisible = true;
        },

        confirmSave() {
            if (this.resumeName.trim() === '') {
                alert('简历名称不能为空');
                return;
            }
            this.saveConfig(this.resumeName);
            this.dialogVisible = false;
        },

        handleClose() {
            this.dialogVisible = false;
        },
        saveResumeData() {
            const { is_display, dialogVisible, ...resumeData } = this.$data; 
            localStorage.setItem(`resumeData_${this.templateName}`, JSON.stringify(resumeData));
            console.log('保存到本地成功');
        },
        async loadIdeas() {
            const response = await getIdeas();
            this.words = response;
            console.log('获取灵感数据成功', this.words);
        },
        load() {
        if (this.count < this.maxCount) {
            this.count += 2
            }
        },
    },
    name: "ResumeEditor", // 简历编辑页
    beforeRouteLeave(to, from, next) {
        const resumeId = Number(this.$route.query.resume_id);
        console.log('离开路由前', resumeId);
        if (resumeId === -1) {
            this.saveResumeData();
            alert(`已自动保存草稿`);
        } else {
            console.log('来自数据库的简历，不保存到本地');
        }
        next();
    },

    // 刷新或关闭页面时保存数据
    mounted() {
        this.loadIdeas();
        this.count = 10;
        window.addEventListener('beforeunload', this.saveResumeData);
    },
    beforeDestroy() {
        window.removeEventListener('beforeunload', this.saveResumeData);
    }
}
</script>
