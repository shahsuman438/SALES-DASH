import axiosInstance from "../interceptor/axiosConfig"


const getSummeryReport = () => {
    return axiosInstance.get("/reports/summery")
}
const getSalesByProductReports = () => {
    return axiosInstance.get("/reports/sales-by-product")
}

const reportService = {
    getSummeryReport,
    getSalesByProductReports
}

export default reportService