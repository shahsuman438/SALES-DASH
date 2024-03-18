import { useEffect, useState } from "react";
import { salesByProductType } from "../../types/types";
import reportService from "../../services/report.service";
import Table from "../../components/table/Table";

const SalesByProduct = () => {
    const [data, setData] = useState<salesByProductType[]>([]);

    useEffect(() => {
        reportService.getSalesByProductReports()
            .then(res => {
                setData(res.data)
            })
            .catch(err => {
                console.log(err)
            })
    }, [])
    return (
        <div>
            <h1 className='text-blue'>Report Sales By Product</h1>
            <Table data={data} />
        </div>
    )
}

export default SalesByProduct