curl --location --request POST 'http://0.0.0.0:3000/generateReport' \
--output 'report.pdf' \
--header 'Content-Type: application/json' \
--data-raw '{
    "nmap": [
        {
            "nmaprun": {
                "scanner": "nmap",
                "args": "nmap -v -A -p- -oG dg.txt -oX dx.txt -oN dl.txt -iL ipList.txt",
                "start": "1618789781",
                "startstr": "Mon Apr 19 06:49:41 2021",
                "version": "7.80",
                "xmloutputversion": "1.04",
                "host": [
                    {
                        "starttime": "1618789781",
                        "endtime": "1618790933",
                        "status": {
                            "state": "up",
                            "reason": "syn-ack",
                            "reason_ttl": "0"
                        },
                        "address": {
                            "addr": "159.203.96.214",
                            "addrtype": "ipv4"
                        },
                        "hostnames": "\n",
                        "ports": {
                            "extraports": {
                                "state": "closed",
                                "count": "65528",
                                "extrareasons": {
                                    "reason": "conn-refused",
                                    "count": "65528"
                                }
                            },
                            "port": [
                                {
                                    "protocol": "tcp",
                                    "portid": "22",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "ssh",
                                        "product": "OpenSSH",
                                        "version": "8.2p1 Ubuntu 4ubuntu0.1",
                                        "extrainfo": "Ubuntu Linux; protocol 2.0",
                                        "ostype": "Linux",
                                        "method": "probed",
                                        "conf": "10",
                                        "cpe": [
                                            "cpe:/a:openbsd:openssh:8.2p1",
                                            "cpe:/o:linux:linux_kernel"
                                        ]
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "25",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "smtp",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "80",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "http",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "303",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "http",
                                        "product": "Golang net/http server",
                                        "extrainfo": "Go-IPFS json-rpc or InfluxDB API",
                                        "method": "probed",
                                        "conf": "10",
                                        "cpe": "cpe:/a:protocol_labs:go-ipfs"
                                    },
                                    "script": {
                                        "id": "http-title",
                                        "output": "Site doesn'\''t have a title (text/plain; charset=utf-8)."
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "8291",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "unknown",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "8728",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": null
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "11211",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "memcache",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                }
                            ]
                        },
                        "times": {
                            "srtt": "289130",
                            "rttvar": "6489",
                            "to": "315086"
                        }
                    },
                    {
                        "starttime": "1618789781",
                        "endtime": "1618790933",
                        "status": {
                            "state": "up",
                            "reason": "syn-ack",
                            "reason_ttl": "0"
                        },
                        "address": {
                            "addr": "165.227.186.93",
                            "addrtype": "ipv4"
                        },
                        "hostnames": "\n",
                        "ports": {
                            "extraports": {
                                "state": "closed",
                                "count": "65528",
                                "extrareasons": {
                                    "reason": "conn-refused",
                                    "count": "65528"
                                }
                            },
                            "port": [
                                {
                                    "protocol": "tcp",
                                    "portid": "22",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "ssh",
                                        "product": "OpenSSH",
                                        "version": "8.2p1 Ubuntu 4ubuntu0.1",
                                        "extrainfo": "Ubuntu Linux; protocol 2.0",
                                        "ostype": "Linux",
                                        "method": "probed",
                                        "conf": "10",
                                        "cpe": [
                                            "cpe:/a:openbsd:openssh:8.2p1",
                                            "cpe:/o:linux:linux_kernel"
                                        ]
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "25",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "smtp",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "80",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "http",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "303",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "http",
                                        "product": "Golang net/http server",
                                        "extrainfo": "Go-IPFS json-rpc or InfluxDB API",
                                        "method": "probed",
                                        "conf": "10",
                                        "cpe": "cpe:/a:protocol_labs:go-ipfs"
                                    },
                                    "script": {
                                        "id": "http-title",
                                        "output": "Site doesn'\''t have a title (text/plain; charset=utf-8)."
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "8291",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "unknown",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "8728",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": null
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "11211",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "memcache",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                }
                            ]
                        },
                        "times": {
                            "srtt": "290287",
                            "rttvar": "3840",
                            "to": "305647"
                        }
                    }
                ],
                "runstats": {
                    "finished": {
                        "time": "1618790933",
                        "timestr": "Mon Apr 19 07:08:53 2021",
                        "elapsed": "1152.40",
                        "summary": "Nmap done at Mon Apr 19 07:08:53 2021; 2 IP addresses (2 hosts up) scanned in 1152.40 seconds",
                        "exit": "success"
                    },
                    "hosts": {
                        "up": "2",
                        "down": "0",
                        "total": "2"
                    }
                },
                "hostport": [
                    {
                        "port": "22"
                    },
                    {
                        "port": "25"
                    },
                    {
                        "port": "80"
                    },
                    {
                        "port": "303"
                    },
                    {
                        "port": "8291"
                    },
                    {
                        "port": "8728"
                    },
                    {
                        "port": "11211"
                    }
                ],
                "hostservice": []
            }
        },
        {
            "nmaprun": {
                "scanner": "nmap",
                "args": "nmap -v -A -p- -oG dg.txt -oX dx.txt -oN dl.txt -iL ipList.txt",
                "start": "1618789781",
                "startstr": "Mon Apr 19 06:49:41 2021",
                "version": "7.80",
                "xmloutputversion": "1.04",
                "host": [
                    {
                        "starttime": "1618789781",
                        "endtime": "1618790933",
                        "status": {
                            "state": "up",
                            "reason": "syn-ack",
                            "reason_ttl": "0"
                        },
                        "address": {
                            "addr": "159.203.96.214",
                            "addrtype": "ipv4"
                        },
                        "hostnames": "\n",
                        "ports": {
                            "extraports": {
                                "state": "closed",
                                "count": "65528",
                                "extrareasons": {
                                    "reason": "conn-refused",
                                    "count": "65528"
                                }
                            },
                            "port": [
                                {
                                    "protocol": "tcp",
                                    "portid": "22",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "ssh",
                                        "product": "OpenSSH",
                                        "version": "8.2p1 Ubuntu 4ubuntu0.1",
                                        "extrainfo": "Ubuntu Linux; protocol 2.0",
                                        "ostype": "Linux",
                                        "method": "probed",
                                        "conf": "10",
                                        "cpe": [
                                            "cpe:/a:openbsd:openssh:8.2p1",
                                            "cpe:/o:linux:linux_kernel"
                                        ]
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "25",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "smtp",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "80",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "http",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "303",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "http",
                                        "product": "Golang net/http server",
                                        "extrainfo": "Go-IPFS json-rpc or InfluxDB API",
                                        "method": "probed",
                                        "conf": "10",
                                        "cpe": "cpe:/a:protocol_labs:go-ipfs"
                                    },
                                    "script": {
                                        "id": "http-title",
                                        "output": "Site doesn'\''t have a title (text/plain; charset=utf-8)."
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "8291",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "unknown",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "8728",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": null
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "11211",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "memcache",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                }
                            ]
                        },
                        "times": {
                            "srtt": "289130",
                            "rttvar": "6489",
                            "to": "315086"
                        }
                    },
                    {
                        "starttime": "1618789781",
                        "endtime": "1618790933",
                        "status": {
                            "state": "up",
                            "reason": "syn-ack",
                            "reason_ttl": "0"
                        },
                        "address": {
                            "addr": "165.227.186.93",
                            "addrtype": "ipv4"
                        },
                        "hostnames": "\n",
                        "ports": {
                            "extraports": {
                                "state": "closed",
                                "count": "65528",
                                "extrareasons": {
                                    "reason": "conn-refused",
                                    "count": "65528"
                                }
                            },
                            "port": [
                                {
                                    "protocol": "tcp",
                                    "portid": "22",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "ssh",
                                        "product": "OpenSSH",
                                        "version": "8.2p1 Ubuntu 4ubuntu0.1",
                                        "extrainfo": "Ubuntu Linux; protocol 2.0",
                                        "ostype": "Linux",
                                        "method": "probed",
                                        "conf": "10",
                                        "cpe": [
                                            "cpe:/a:openbsd:openssh:8.2p1",
                                            "cpe:/o:linux:linux_kernel"
                                        ]
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "25",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "smtp",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "80",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "http",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "303",
                                    "state": {
                                        "state": "open",
                                        "reason": "syn-ack",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "http",
                                        "product": "Golang net/http server",
                                        "extrainfo": "Go-IPFS json-rpc or InfluxDB API",
                                        "method": "probed",
                                        "conf": "10",
                                        "cpe": "cpe:/a:protocol_labs:go-ipfs"
                                    },
                                    "script": {
                                        "id": "http-title",
                                        "output": "Site doesn'\''t have a title (text/plain; charset=utf-8)."
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "8291",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "unknown",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "8728",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": null
                                    }
                                },
                                {
                                    "protocol": "tcp",
                                    "portid": "11211",
                                    "state": {
                                        "state": "filtered",
                                        "reason": "no-response",
                                        "reason_ttl": "0"
                                    },
                                    "service": {
                                        "name": "memcache",
                                        "method": "table",
                                        "conf": "3"
                                    }
                                }
                            ]
                        },
                        "times": {
                            "srtt": "290287",
                            "rttvar": "3840",
                            "to": "305647"
                        }
                    }
                ],
                "runstats": {
                    "finished": {
                        "time": "1618790933",
                        "timestr": "Mon Apr 19 07:08:53 2021",
                        "elapsed": "1152.40",
                        "summary": "Nmap done at Mon Apr 19 07:08:53 2021; 2 IP addresses (2 hosts up) scanned in 1152.40 seconds",
                        "exit": "success"
                    },
                    "hosts": {
                        "up": "2",
                        "down": "0",
                        "total": "2"
                    }
                },
                "hostport": [
                    {
                        "port": "22"
                    },
                    {
                        "port": "25"
                    },
                    {
                        "port": "80"
                    },
                    {
                        "port": "303"
                    },
                    {
                        "port": "8291"
                    },
                    {
                        "port": "8728"
                    },
                    {
                        "port": "11211"
                    }
                ],
                "hostservice": []
            }
        }
    ]
}'
