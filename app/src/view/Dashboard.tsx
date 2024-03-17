import Card from '../components/card/Card'

const Dashboard = () => {
    return (
        <div className="d-flex justify-content-space-between pt-5 p-5 ">
            <Card title="Most Profitable Product" value="Computer" />
            <Card title="Least Profitable Product" value="Speaker" />
            <Card title="Date of Highest Sales" value="2024-5-5" />
            <Card title="Date of Least Sales" value="2024-4-4" />
        </div>
    )
}

export default Dashboard;