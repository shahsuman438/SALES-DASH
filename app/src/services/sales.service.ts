import axiosInstance from '../interceptor/axiosConfig';

const getAllSales = () => {
  return axiosInstance.get('/sales');
};

const salesService = {
  getAllSales,
};

export default salesService;
