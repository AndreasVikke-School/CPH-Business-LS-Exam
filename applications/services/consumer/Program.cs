using Grpc.Net.Client;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

IHostBuilder host = Host.CreateDefaultBuilder(args)
    .ConfigureServices((hostContext, services) =>
    {
        services.AddOptions<KafkaSettings>().BindConfiguration("Kafka");
        services.AddSingleton<KafkaConsumer>();
        services.AddSingleton<RedisConsumer>();
        services.AddSingleton<PostgresConsumer>();

        RedisSettings redisSettings = hostContext.Configuration.GetSection("Redis").Get<RedisSettings>();
        services.AddSingleton(services => new AttendanceCodeProto.AttendanceCodeProtoClient(GrpcChannel.ForAddress(redisSettings.ServiceUrl)));

        PostgresSettings postgresSettings = hostContext.Configuration.GetSection("Postgres").Get<PostgresSettings>();
        services.AddSingleton(services => new CheckInService.CheckInServiceClient(GrpcChannel.ForAddress(postgresSettings.ServiceUrl)));

        services.AddHostedService<ConsumerService>();
    });

host.Build().Run();

public class ConsumerService : BackgroundService
{
    private readonly KafkaConsumer _kafkaConsumer;
    private readonly RedisConsumer _redisConsumer;
    private readonly PostgresConsumer _postgresConsumer;
    private readonly ILogger<ConsumerService> _logger;

    public ConsumerService(KafkaConsumer kafkaConsumer, RedisConsumer redisConsumer, PostgresConsumer postgresConsumer, ILogger<ConsumerService> logger)
    {
        _kafkaConsumer = kafkaConsumer;
        _redisConsumer = redisConsumer;
        _postgresConsumer = postgresConsumer;
        _logger = logger;
    }

    protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        => await Task.Run(async () 
            => {
                _logger.LogInformation("Started polling kafka broker for new attendance events");

                while (!stoppingToken.IsCancellationRequested)
                {
                    AttendanceEvent attendance = _kafkaConsumer.Poll(stoppingToken);
                    CodeValidity codeValidity = _redisConsumer.CheckCodeValidity(attendance);
                    await _postgresConsumer.SaveAttendanceEvent(attendance, codeValidity);
                }
            }, stoppingToken);
}