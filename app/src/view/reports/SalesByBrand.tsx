import { useEffect, useState } from "react";
import { salesByProductType } from "../../types/types";
import reportService from "../../services/report.service";
import Table from "../../components/table/Table";

const SalesByBrand = () => {
    const [data, setData] = useState<salesByProductType[]>([]);

    useEffect(() => {
        reportService.getSalesByBrandReports()
            .then(res => {
                setData(res.data)
            })
            .catch(err => {
                console.log(err)
            })
    }, [])
    return (
        <div>
            <h1 className='text-blue'>Report Sales By Brand</h1>
            <Table data={data} itemsPerPage={100} />
        </div>
    )
}

export default SalesByBrand