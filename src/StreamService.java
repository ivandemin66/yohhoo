import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;
import stream.StreamServiceGrpc;
import stream.StreamRequest;
import stream.StreamResponse;

import java.io.IOException;

public class StreamService extends StreamServiceGrpc.StreamServiceImplBase {
    @Override
    public void startStream(StreamRequest request, StreamObserver<StreamResponse> responseObserver) {
        String message = "Stream started for user: " + request.getUserId();
        StreamResponse response = StreamResponse.newBuilder().setStatus("SUCCESS").setMessage(message).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void endStream(StreamRequest request, StreamObserver<StreamResponse> responseObserver) {
        String message = "Stream ended for user: " + request.getUserId();
        StreamResponse response = StreamResponse.newBuilder().setStatus("SUCCESS").setMessage(message).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    public static void main(String[] args) throws IOException, InterruptedException {
        Server server = ServerBuilder.forPort(50052)
                .addService(new StreamService())
                .build();

        System.out.println("Stream Service is running on port 50052");
        server.start();
        server.awaitTermination();
    }
}
