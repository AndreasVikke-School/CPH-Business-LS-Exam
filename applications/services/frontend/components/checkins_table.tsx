type tplotOptions = {
    [key: string]: string
}

const statuses: tplotOptions = {
    "1": "success",
    "2": "outOfTime",
    "3": "notFound",
    "4": "outOfRange",
    "5": "error"
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
                {"checkIn" in data ? data.checkIn.sort((a: any, b: any) => (a.checkinTime < b.checkinTime) ? 1 : -1).map((d: any) => {
                    var c = d.status == 1 ? "success" : d.status == 2 || d.status == 4 ? "warning" : "danger"
                    var date = new Date(d.checkinTime)
                    return (
                        <tr key={d.checkinTime}>
                            <th scope="row">{d.attendanceCode}</th>
                            <td>{date.toDateString()} - {date.getHours()}:{date.getMinutes()}</td>
                            <td><span className={"badge bg-" + c}>{statuses[d.status]}</span></td>
                        </tr>
                    )
                }) : null}
            </tbody>
        </table>
    )
}
export default CheckInTable