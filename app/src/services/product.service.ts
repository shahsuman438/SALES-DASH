import axiosInstance from '../interceptor/axiosConfig';

const getAllProducts = () => {
  return axiosInstance.get('/product');
};

const productService = {
  getAllProducts,
};

export default productService;
