using Confluent.Kafka;
using System.Text.Json;

namespace LSExam.Services;

public class KafkaConsumer
{
    private readonly ILogger _logger;
    private readonly string _topic;
    private readonly IConsumer<Null, string> _consumer;

    public KafkaConsumer(ILogger<KafkaConsumer> logger, IOptions<KafkaSettings> options)
    {
        _logger = logger;
        KafkaSettings option = options.Value;

        ConsumerConfig config = new()
        {
            BootstrapServers = option.kafkaBrokers,
            GroupId = option.GroupId,
            AllowAutoCreateTopics = true,
            AutoOffsetReset = AutoOffsetReset.Earliest,
            SecurityProtocol = SecurityProtocol.Plaintext,
            EnableAutoCommit = false,
        };
        _topic = option.CheckinTopic;

        _consumer = new ConsumerBuilder<Null, string>(config).Build();
        _consumer.Subscribe(_topic);
        _consumer.Assign(new TopicPartition(_topic, Partition.Any));
    }

    public AttendanceEvent Poll(CancellationToken ct)
    {
        var result = _consumer.Consume(ct);
        var message = result.Message;

        _consumer.Commit();

        AttendanceEvent value = JsonSerializer.Deserialize<AttendanceEvent>(message.Value) ?? new();
        _logger.LogInformation($"Consumed attandance event: {value}");

        return value;
    }
}

