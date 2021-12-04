namespace LSExam.Configs;

public class KafkaSettings
{
    public string kafkaBrokers { get; set; } = "";
    public string CheckinTopic { get; set; } = "";
    public string GroupId { get; set; } = "";
}
