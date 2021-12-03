using Confluent.Kafka;
using System.Text.Json;

namespace LSExam.Services;

public class KafkaConsumer
{
    private readonly ILogger _logger;
    private readonly ConsumerBuilder<Null, string> _builder;
    private readonly string _topic;

    public KafkaConsumer(ILogger<KafkaConsumer> logger, IOptions<KafkaSettings> options)
    {
        _logger = logger;
        KafkaSettings option = options.Value;

        ConsumerConfig config = new()
        {
            BootstrapServers = option.kafkaBrokers,
            GroupId = "consumer",
            AllowAutoCreateTopics = true,
            AutoOffsetReset = AutoOffsetReset.Earliest,
            SecurityProtocol = SecurityProtocol.Plaintext,
            EnableAutoCommit = false
        };
        _topic = option.CheckinTopic;

        _builder = new ConsumerBuilder<Null, string>(config);
    }

    public AttendanceEvent Poll(CancellationToken ct)
    {
        using IConsumer<Null, string> consumer = _builder.Build();
        consumer.Subscribe(_topic);
        consumer.Assign(new TopicPartition(_topic, Partition.Any));

        var result = consumer.Consume(ct);
        var message = result.Message;

        consumer.Commit(result);

        AttendanceEvent value = JsonSerializer.Deserialize<AttendanceEvent>(message.Value) ?? new();
        _logger.LogInformation($"Consumed attandance event: {value}");

        return value;
    }
}

