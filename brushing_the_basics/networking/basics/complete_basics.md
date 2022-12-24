# Computer Networks

## Client Server Architecture

- A **client** is a machine that requests information from another machine which is called **server**

## Protocols

- TCP (Transmission Control Protocol)
  - It ensures data is transferred uncorupped

- UDP (User Datagram Protocol)
  - If we don't care, if we recieve 100% of the data. Example - Video Conference data

- HTTP (HyperText Transfer Protocol)
  - It defines the format in which data will be consumed

## Addresses

- Computers on the internet are identified by IP Addresses
- Every domains. such as `google.com` are linked to an IP address
- Format of an IP Address: `X.X.X.X`; where `X` is a number between 0 and 255
- Some IP addresses are reserved
- To know your IP address (on Ubuntu): `curl ifconfig.me -s`

## ISPs and Modems

- ISP (Internet Service Providers) are companies that provides you devices to access the Internet. Example - AT&T, Vodafone
- They provide a Modem/Router which has a global IP address.
- The modem assigns IP to any device which gets connected to the modem. These are known as local IP addresses. This assignment of IP is done through DHCP (Dynamic Host Control Protocol)
- If we access any website from any one of the connected device, that website will only be able to see the global IP address
- Once the response is recieved from the website, the modem will use NAT (Network Access Translator) and see which device it has to send the response

## Ports

- As we saw above, the router can "route" the response from the website, to the specific device that made a request to that website.
- What if we want to send that response to a particular application, which made that request and is present on the device that response was routed to?
- This is done using **Ports**
- Basically, every application (on a device) runs on same IP address. **Ports** helps us to create a distinction between them
- Example: `172.12.3.3:3456`; where `3456` is the port number.
- These are 16-bit numbers. The total number of possible Ports are 65536.
- Well known Ports:
  - HTTP - 80
  - HTTPS - 443
- Reserverd Ports are between 0 and 1023

## Network Connection Types

- LAN 
  - It's restricted to a particular area
  - Ethernets, WIFI

- MAN (Metropolitan Area Network)
  - Accross the city

- WAN (Wide Area Network)
  - Accross countries
  - Optical Fiber cables

- All of these come together and form the internet.

## Network Topologies

- Bus Topology
  - Single Bus, all the computers are connected to it

- Ring Topology
  - All connected in a ring shape
  - Sending a message between two computers require the data to go through computers between these two

- Start Topology
  - One central device connected to all computers
  - If central device fails, the complete network breaks

- Tree Topology
  - Bus +_Start

- Mesh Topology
  - Every Single computer is connected to every single computer
  - Pretty expensive


## OSI (Open Systems Interconnection) Model

- Application Layer
  - It's the Software that users interact with

- Presentation Layer
  - It recieves the data from the application layer and converts them into binary format
  - It encrypts the data
  
- Session Layer
  - The session layer is responsible for establishing, maintaining, and terminating communication sessions between devices on a network.
  - It allows two devices to communicate with each other in a full-duplex mode
  - RPC and NetBIOS protocols operate at this level

- Transport Layer
  - Responsible for the transfer of data
    - Data recieved from Session layer, it is divided into smaller data units called segments
    - These segments will contain the source and destination port number, sequence number
    - TCP is connection-oriented transmission and UDP is connectionless-oriented transmission 

- Network Layer
  - The network layer is responsible for providing logical addressing and routing services for data transmission between devices on a network. It assigns unique logical addresses (also known as IP addresses) to each device on a network and uses these addresses to route data packets between devices
  - Protocols that operate here: IP, ICMP, IGMP


- Data Link Layer
  - Physical addressing (MAC addresses) is done here
  - MAC address is the address of the network interface of the PC

- Physical Layer
  - Hardware Layer
  - Transport 0s and 1s through cables