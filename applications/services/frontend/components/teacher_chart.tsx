import { Bar, getElementAtEvent } from 'react-chartjs-2';
import { Chart, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend } from 'chart.js'
import { useEffect, useRef, useState } from 'react';
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


const TeacherChart = ({ checkins }: any) => {
    const [data, setData] = useState({labels: [] as any, datasets: [{} as any]} as any)
    const chartRef = useRef();
    checkins = "checkIn" in checkins ? checkins : { checkIn: [] }

    const labels = Array.from(new Set(checkins.checkIn.map((d: any) => d.attendanceCode)))

    useEffect(() => {
        setData({
            labels: labels,
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
        })
    }, [checkins])

    const click = (event: any) => {
        if (chartRef.current) {
            var activePoints = getElementAtEvent(chartRef.current, event)
            if (activePoints.length > 0) {
                var attCode = labels[activePoints[0].index]
                var label = labels.filter((l: any) => l == attCode)
                setData({
                    labels: label,
                    datasets: [
                        {
                            label: 'Success',
                            data: label.map((l: any) => checkins.checkIn.filter((d: any) => d.attendanceCode == l && d.status == 1).length),
                            backgroundColor: 'rgba(10, 255, 19, 0.5)',
                        },
                        {
                            label: 'Out Of Time',
                            data: label.map((l: any) => checkins.checkIn.filter((d: any) => d.attendanceCode == l && d.status == 2).length),
                            backgroundColor: 'rgba(255, 255, 19, 0.5)',
                        },
                        {
                            label: 'Out Of Range',
                            data: label.map((l: any) => checkins.checkIn.filter((d: any) => d.attendanceCode == l && d.status == 4).length),
                            backgroundColor: 'rgba(255, 255, 19, 0.5)',
                        },
                        {
                            label: 'Error',
                            data: label.map((l: any) => checkins.checkIn.filter((d: any) => d.attendanceCode == l && d.status == 4).length),
                            backgroundColor: 'rgba(255, 19, 19, 0.5)',
                        },
                    ],
                })
            }
        }
    }

    return (
        <Bar ref={chartRef} options={options} data={data} onClick={click} />
    )
}
export default TeacherChart