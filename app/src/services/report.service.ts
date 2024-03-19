import axiosInstance from '../interceptor/axiosConfig';

const getSummeryReport = () => {
  return axiosInstance.get('/reports/summery');
};

const getSalesByBrandReports = () => {
  return axiosInstance.get('/reports/sales-by-brand');
};

const getSalesByProductReports = () => {
  return axiosInstance.get('/reports/sales-by-product');
};

const reportService = {
  getSummeryReport,
  getSalesByProductReports,
  getSalesByBrandReports,
};

export default reportService;
