type tplotOptions = {
    [key: string]: string
}

const statuses: tplotOptions = {
    "0": "success",
    "1": "outOfTime",
    "2": "notFound",
    "3": "outOfRange",
    "4": "error"
}

const CheckInTable = ({ data }: { data: any }) => {
    return (
        <table className="table">
            <thead>
                <tr>
                    <th key="1" scope="col">Attendance Code</th>
                    <th key="2" scope="col">CheckIn Time</th>
                    <th key="3" scope="col">Status</th>
                </tr>
            </thead>
            <tbody>
                {"checkIn" in data ? data.checkIn.map((d:any) => {
                    var c = d.status == 0 ? "success" : d.status == 1 || d.status == 3 ? "warning" : "danger"
                    var date = new Date(d.checkinTime)
                    return (
                        <tr key={d.attendanceCode}>
                            <th scope="row">{d.attendanceCode}</th>
                            <td>{date.toDateString()} - {date.getHours()}:{date.getMinutes()}</td>
                            <td><span className={"badge bg-" + c}>{statuses[d.status]}</span>{d.status}</td>
                        </tr>
                    )
                }) : ""}
            </tbody>
        </table>
    )
}
export default CheckInTable