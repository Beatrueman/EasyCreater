export interface Response {
    msg: string;
    status: number;
}

export interface ResumeData {
    resume_id: number;
    username: string;
    user_id: number;
    template_name: string;
    resume_data: string;
    Timestamp: string;
    IsShared: boolean;
  }