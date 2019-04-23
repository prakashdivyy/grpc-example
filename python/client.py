import grpc
import example_pb2 as pb
import example_pb2_grpc


def main():

    # Create channel and stub to server's address and port.
    channel = grpc.insecure_channel('localhost:50050')
    stub = example_pb2_grpc.ExampleServiceStub(channel)

    # Exception handling.
    try:
        response = stub.SayHello(pb.HelloRequest(greeting="Prakash"))
        print(response)

    # Catch any raised errors by grpc.
    except grpc.RpcError as e:
        print("Error raised: " + e.details())

    # Exception handling.
    try:
        response = stub.DoAddition(
            pb.NumberRequest(first_number=1, second_number=2))
        print(response)

    # Catch any raised errors by grpc.
    except grpc.RpcError as e:
        print("Error raised: " + e.details())

    # Exception handling.
    try:
        response = stub.DoSubtraction(
            pb.NumberRequest(first_number=1, second_number=2))
        print(response)

    # Catch any raised errors by grpc.
    except grpc.RpcError as e:
        print("Error raised: " + e.details())


if __name__ == '__main__':
    main()
