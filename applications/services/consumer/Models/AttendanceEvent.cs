using System.Text.Json.Serialization;

namespace LSExam.Models;
public record AttendanceEvent
{
    [JsonPropertyName("attendanceCode")]
    public long AttendanceCode { get; set; } = 0;
    [JsonPropertyName("studentId")]
    public string StudentId { get; set; } = "";
    [JsonPropertyName("currentUnixTime")]
    public long CurrentUnixTime { get; set; } = 0;
    [JsonPropertyName("lat")]
    public double Latitude { get; set; } = 0;
    [JsonPropertyName("long")]
    public double Longitude { get; set; } = 0;

    public AttendanceEvent() { }

    public AttendanceEvent(long attendanceCode, string studentId, long currentUnixTime, double latitude, double longitude)
    {
        AttendanceCode = attendanceCode;
        StudentId = studentId;
        CurrentUnixTime = currentUnixTime;
        Latitude = latitude;
        Longitude = longitude;
    }
}