import axiosInstance from "../interceptor/axiosConfig"


const getSummeryReport = () => {
    return axiosInstance.get("/reports/summery")
}

const reportService = {
    getSummeryReport
}

export default reportService