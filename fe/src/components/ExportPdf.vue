<template>
    <CustomButton btn-type="primary" @click="exportPdf">
        导出为PDF
    </CustomButton>
</template>


<script>
import CustomButton from '../components/CustomButton.vue';

export default {
    components: {
        CustomButton
    },

    pops: {
        resumeFormat: {
            type: String
        },

    }, 

    methods: {
        exportPdf() {
            const unit = this.resumeFormat == 'a4' ? 'mm': 'in';

            const pdfConfig = {
                margin: 0,
                filename: "简历.pdf",
                jsPDF: {
                    unit: unit,
                    format: this.resumeFormat
                } 
            }
            const resume = document.getElementById('resume');
            html2pdf().from(resume).save();
        }
    }

}
</script>