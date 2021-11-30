using Grpc.Net.Client;
using LSExam.Configs;
using LSExam.Services;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Protos;

IHostBuilder host = Host.CreateDefaultBuilder(args)
    .ConfigureServices((hostContext, services) =>
    {
        services.AddOptions<KafkaSettings>().BindConfiguration("Kafka");
        services.AddSingleton<KafkaConsumer>();
        services.AddSingleton<RedisConsumer>();

        RedisSettings redisSettings = hostContext.Configuration.GetSection("Redis").Get<RedisSettings>();
        services.AddSingleton(services => new AttendanceCodeProto.AttendanceCodeProtoClient(GrpcChannel.ForAddress(redisSettings.ServiceUrl)));
        services.AddHostedService<ConsumerService>();
    });

host.Build().Run();

public class ConsumerService : BackgroundService
{
    private readonly KafkaConsumer _kafkaConsumer;
    private readonly RedisConsumer _redisConsumer;

    public ConsumerService(KafkaConsumer kafkaConsumer, RedisConsumer redisConsumer)
    {
        _kafkaConsumer = kafkaConsumer;
        _redisConsumer = redisConsumer;
    }

    protected override Task ExecuteAsync(CancellationToken stoppingToken)
        => Task.Run(() 
            => {
                while (!stoppingToken.IsCancellationRequested)
                {
                    var attendance = _kafkaConsumer.Poll(stoppingToken);
                    var codeValidity = _redisConsumer.CheckCodeValidity(attendance);
                    Console.WriteLine(codeValidity.ToString());
                }
            }, stoppingToken);
}