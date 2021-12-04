import { Bar } from 'react-chartjs-2';
import { Chart, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend } from 'chart.js'
Chart.register(
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
);

export const options = {
    responsive: true,
    plugins: {
        legend: {
            position: 'top' as const,
        },
        title: {
            display: true,
            text: 'CheckIns For Each Attendance Code',
        },
    },
};


const TeacherChart = ({ checkins } : any) => {
    checkins = "checkIn" in checkins ? checkins : { checkIn: []}

    const labels = checkins.checkIn.map((d: any) => d.attendanceCode)

    const data = {
        labels,
        datasets: [
            {
                label: 'Success',
                data: labels.map((l: any) => checkins.checkIn.filter((d: any) => d.attendanceCode == l && d.status == 1).length),
                backgroundColor: 'rgba(10, 255, 19, 0.5)',
            },
            {
                label: 'Other',
                data: labels.map((l: any) => checkins.checkIn.filter((d: any) => d.attendanceCode == l && d.status != 1).length),
                backgroundColor: 'rgba(255, 19, 19, 0.5)',
            },
        ],
    };

    return (
        <Bar options={options} data={data} />
    )
}
export default TeacherChart