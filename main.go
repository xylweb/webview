package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"

	//"io/ioutil"
	"time"

	"fmt"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"github.com/riftbit/go-systray"
	"github.com/webview/webview"
)

var (
	icobase64  = "AAABAAEALi4AAAEAIACoIgAAFgAAACgAAAAuAAAAXAAAAAEAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADCvLj/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/CvLj/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////NP///zT///80////NP///zT///81/Pz/N/Dv/z3Kxf9LfXL/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP39/zzRzP9PZVf/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///80////OeXi/09qXf9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////OOvp/zvV0f871dH/O9XR/znl4v81/Pz/NP///zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///z/Cvf9WQC//VkAv/1ZAL/9WQC//Tm5i/zre2/80////NP///zT///80////N+/u/0x2av9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VkAv/1ZAL/9KgXb/NvTz/zT///80////NP///zX6+f9HlYv/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////P8K9/1ZAL/9WQC//VkAv/1ZAL/9WQC//T2ZZ/zjm5P80////NP///zT///80////RKeg/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///z/Cvf9WQC//VkAv/1ZAL/9WQC//VkAv/09oWv845+X/NP///zT///80////NP39/0WhmP9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VkAv/1ZAL/9DrKT/Nfz8/zT///80////NP///zb08/9KgXb/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////PsnE/0x4bP9MeGz/THhs/0iMgv880cz/Nfv6/zT///80////NP///zT///8+ycT/VEk5/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///863dn/T2pd/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zfx7/9Drab/U1BB/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////NP///zT///80////NP///zT///80////NP39/zva1v9KgXb/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///zva1v9DraX/Qq+o/z3NyP837Or/NP7+/zT///80////Nfz8/0C8tv9USTn/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VEk5/0G2sP80/f3/NP///zT///80/v7/P8C6/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////P8K9/1ZAL/9WQC//VkAv/1ZAL/9PZln/OObk/zT///80////NP///zb08/9KgXb/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///z/Cvf9WQC//VkAv/1ZAL/9WQC//VkAv/zvY1f80////NP///zT///80/v7/RKWd/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VkAv/01wY/846uj/NP///zT///80////NP///0SnoP9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////P8K9/1ZAL/9WQC//VkAv/0t9cv872tf/NP7+/zT///80////NP///zT9/f9FoZj/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///837uz/TXRo/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///81+/r/Q6uk/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///846uj/SYqA/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9IkIb/O9XR/zvV0f871dH/O9XR/zvV0f871dH/O9XR/zvV0f871dH/PNPP/0G0rf9LgHX/VEk5/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/CvLj/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/CvLj/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
	ico2base64 = "AAABAAEAMDAAAAEAIACoJQAAFgAAACgAAAAwAAAAYAAAAAEAIAAAAAAAACQAAAAAAAAAAAAAAAAAAAAAAAD+/v7////////////+/v7///////7+/v///////v7+/////////////v7+///////+/v7//v7+//3+/v/+/v7/+v3+/8rn/P+Jyfn/UrD5/yed+f8Hj/n/AIv6/wCM+/8AjPr/AIv7/wWO+f8jm/j/Ta75/4DG+v/D4/v/+Pz+//7+/v/+/v7//v7+/////////////v7+/////////////v7+/////////////v7+/////////////v7+///////+/v7//v7+//7+/v/+/v7///////7+/v///////v7+/////////////v7+//7+/v/+/v///v7+/9ft+/91wPn/H5j5/wGL+v8AjPr/AIv7/wCL+/8Ai/v/AIv6/wCL+/8Ai/r/AIv6/wCL+v8Ai/v/AIv6/wCL+v8Bi/r/GJX4/2m7+f/N6Pv//f7+//7+/v/+/v7//v7+//7+/v///////v7+/////////////v7+//7+/v///////v7+///////+/v7//v7+//7+/v/+/v7//v7+//7+/v/+/v7//v7+//7+/v/+/v7//v7+//3+/f/R6vv/U7H5/wWN+f8Ai/r/AIv6/wCL+/8Ai/v/AIv6/wCL+v8AjPr/AIv6/wCM+v8Ai/r/AIv6/wCM+v8Ai/r/AIv6/wCL+v8Ai/r/AIv6/wCL+v8CjPr/Rqv5/8Tk+//9/v7//v7+//7+/v/+/v7//v7+//7+/v/+/v7//v7+//7+/v/+/v7//v7+//7+/v////////////7+/v/+/v7///////7+/v///////v7+//7+/v/+/v7/7vj9/228+f8Gjfn/AIv7/wCL+v8Ai/v/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIz6/wCM+/8Ai/r/AIv7/wKM+v9btPn/5/T8//3+/v/+/v7//v7+/////////////v7+/////////////v7+///////////////////////////////////////+//7//v7+//3+/v/D4/v/Ipr4/wCM+v8Ai/r/AIz7/wCM+/8Ai/v/AIz7/wCM+/8AjPv/AIz7/wCM+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCM+/8Ai/v/AIz7/wCM+/8AjPv/AIz7/wCM+/8Ai/r/F5X5/7Hb+v/+/v7//v7+///////////////////////////////////////////////////////////////////////+/v7//P7+/5fP+f8Hjvn/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+v8AjPv/AIz6/wCM+/8AjPv/AIz6/wSM+f+Axfn/+/3+/////////////v7+/////////////v7+//////////////////7+/v/+/v7///////3+/v/8/v3/fsT5/wGM+f8Ai/v/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+v8AjPv/AIv6/wCL+v8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+v8AjPv/AIz6/wCL+/8AjPv/AIv6/wCL+/8Bi/r/Z7n4//r8/f///////v7+//7+/v///////v7+///////+/v7///////7+/v/+/v7//v7+//3+/v+Axfj/AYv6/wCL+v8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/r/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/r/AIv6/wCL+/8Ai/v/AIv6/wCL+v8Ai/v/AIv6/wCL+/8Ai/r/AYv7/2e5+P/7/f7//v7+/////////////v7+///////+/v7//v7+//7+/v/+/v7//f7+/5rR+f8Ci/r/AIv7/wCM+v8AjPv/AIz6/wCM+v8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+v8AjPr/AIz6/wCM+/8AjPv/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIv6/wGL+/+Bxfn//f7+//7+/v///////v7+///////+/v7////////////+/v7/yOb8/wiP+f8AjPv/AIz7/wCL+v8AjPv/AIz7/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCL+/8Ejfn/sdv6//3+/v///////v7+//////////////////7+/v/y+f3/J534/wCM+v8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+v8Ai/r/F5X5/+f0/P/9/v7//v7+//////////////7+//7+/v93wfn/AIv7/wCM+v8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIv6/1u1+f/9/v7//v7+///////+/v7//v7+/9nu/P8Jj/n/AIv7/wCM+v8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCL+v8AjPv/AIz6/wCL+v8AjPv/AIz6/wCL+v8AjPv/AIz6/wCL+v8AjPv/AIz6/wCL+v8AjPv/AIz6/wCL+v8AjPv/AIz6/wCL+v8AjPv/AIz6/wOM+v/E5fv//v7+/////////////v79/2K3+f8Ai/v/AIv6/wCL+/8AjPv/AIv7/wmP+v+s2fn/jcr4/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/6za+P8Tk/n/AIz6/wCM+/8AjPv/AIv7/wCL+/9Gq/n//f7+//7+/v//////4/P8/wmO+v8AjPr/AIz7/wCM+v8AjPv/AIz6/zSj+P+Cxvf/kc35/wuP+v8Ai/r/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+/8Ai/v/AIv7/wCL+v8Ci/r/d8H4/4nJ+P9QsPn/AIv6/wCM+/8AjPv/AIz6/wCL+v8CjPr/zej7//7+/v/+/v7/hsj5/wCL+v8AjPr/AIz7/wCM+v8AjPv/AIz6/zak+P9SsPn/Ipr5/5fQ+f8Rk/n/AIz6/wCL+v8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIz6/wCM+/8AjPr/AIv7/wWN+f+Ex/j/Qaj4/zak+v9Rsfn/AIv7/wCM+/8AjPv/AIv6/wCM+/8Ai/r/abv5//3+/v/8/f3/LqD4/wCL+v8Ai/r/AIv7/wCL+v8Ai/v/AIv6/zak+P9RsPn/AIv6/xeW+P+Z0Pj/Gpf5/wCL+v8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+v8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/r/CY/5/47M+P8zovj/AYv6/zWk+v9RsPn/AIv7/wCL+/8Ai/v/AIv6/wCL+v8Ai/v/GJX4//j8/v/b7/z/Aoz6/wCM+v8AjPr/AIz7/wCL+v8AjPr/AIz6/zak+P9RsPn/AIz6/wCL+/8Okvn/lc/4/yWc+P8Ai/r/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8Pkfn/lM/5/yac+P8Ai/v/AIz6/zWk+v9RsPn/AIv7/wCM+v8AjPr/AIz6/wCM+v8Ai/v/AYv7/8Pj+/+d0vr/AIv7/wCM+v8AjPr/AIz7/wCM+v8AjPv/AIz6/zak+P9RsPn/AIz6/wCM+/8Ai/r/CI/6/47M+f8yovj/AIv6/wCM+/8AjPv/AIz6/wCL+/8AjPv/AIz7/wCM+/8Ai/v/AIz6/wCM+/8Ai/v/AIz6/wCM+/8Ai/v/AIv6/xiV+f+Z0fj/G5f5/wGL+/8Ai/r/AIv6/zWk+v9RsPn/AIv7/wCM+/8Ai/v/AIz6/wCM+/8Ai/v/AIv6/4HG+f9ou/n/AIv7/wCL+/8Ai/r/AIz7/wCM+/8Ai/v/AIv7/zak+P9RsPn/AIv7/wCL+/8Ai/v/AIv6/wSN+f+Ex/n/Qqj4/wCM+/8Ai/v/AIv7/wCL+/8Ai/v/AIv6/wCM+v8Ai/r/AIv7/wCL+/8Ai/v/AIv7/wCL+v8Ai/r/IZr5/5bQ+f8Skvn/AYv6/wCL+/8Ai/v/AIv6/zWk+v9RsPn/AIv7/wCL+/8AjPv/AIv7/wCL+/8AjPv/AIv7/02u+f8/qPn/AIv6/wCM+/8AjPr/AIz7/wCM+v8AjPv/AIz6/zak+P9RsPn/AIz6/wCM+/8AjPv/AIv7/wCL+/8CjPr/dsD5/1Kw+f8Ai/r/AIv7/wCM+/8AjPv/AIv6/yue+f9Jq/j/AYv6/wCL+/8AjPv/AIz6/wCL+v8vn/j/kc34/wqQ+v8Ai/v/AIv7/wCM+/8AjPv/AIv6/zWk+v9RsPn/AIv7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz6/yOb+P8gmvn/AIv7/wCM+v8AjPr/AIz7/wCM+v8AjPv/AIz6/zak+P9RsPn/AIz6/wCM+/8AjPv/AIz6/wCM+/8Ai/r/AYv6/2W5+P9juPj/AYv6/wCL+v8Ai/r/OaT5/4rK+f9svPf/YLb4/wGL+v8Ai/r/AYv6/zym+P+Hyfn/BY35/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIv6/zWk+v9RsPn/AIv7/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIv6/wWO+f8MkPn/AIz6/wCL+v8Ai/r/AIv7/wCL+v8Ai/v/AIv6/zak+P9RsPn/AIv6/wCL+v8Ai/v/AIv6/wCL+v8Ai/v/AIv6/wGL+v9Vsfn/dL/4/wKL+v9JrPj/fcT4/wOM+f8Bi/r/WLP5/3G++P8Ci/r/Ta35/3rD+P8DjPn/AIv6/wCL+/8Ai/r/AIv6/wCL+/8Ai/v/AIv6/zWk+v9RsPn/AIv7/wCL+/8Ai/v/AIv6/wCL+v8Ai/v/AIv6/wCM+v8CjPr/AIv6/wCM+v8AjPr/AIz7/wCM+v8AjPv/AIz6/zak+P9RsPn/AIz6/wCM+/8AjPr/AIz6/wCM+/8AjPv/AIz6/wCM+/8Bi/r/RKr3/6XW9/9vvfj/AYz6/wCM+/8AjPr/AYv6/0is9/+k1fj/bb33/wGM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+/8AjPr/AIv6/zWk+v9RsPn/AIv7/wCM+/8AjPr/AIz6/wCM+v8AjPv/AIz6/wCM+/8CjPr/AIv7/wCL+v8Ai/r/AIz7/wCM+v8Ai/v/AIz6/zak+P9RsPn/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8BjPr/arv4/162+f8Bi/r/AIv6/wCM+v8AjPv/AIv7/wCL+/85pPj/i8r5/weO+f8Ai/v/AIz6/wCM+v8AjPv/AIz6/wCM+v8AjPv/AIv6/zWk+v9RsPn/AIv7/wCM+v8AjPv/AIz6/wCM+v8AjPv/AIv6/wCM+/8Okfj/AIz6/wCL+v8Ai/r/AIv7/wCL+v8Ai/v/AIv6/zak+P9RsPn/AIv6/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv7/wKM+v97wvj/Ta74/wCL+v8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+v8Bi/v/K575/5PO+P8MkPn/AIv7/wCL+/8Ai/v/AIv6/wCL+/8Ai/v/AIv6/zWk+v9RsPn/AIv7/wCL+/8Ai/v/AIv6/wCL+/8AjPv/AIv6/wCL+/8jm/j/AIz6/wCL+v8AjPr/AIz7/wCL+v8AjPr/AIz6/zak+P9RsPn/AIz6/wCL+v8Ai/r/AIv6/wCL+v8Ai/r/BY35/4jJ+f89p/j/AIv6/wCM+v8AjPr/AIv6/wCM+v8AjPr/AIv6/wCM+v8Ai/v/AIv6/x+Y+f+X0fn/FJT5/wCL+v8Ai/r/AIv6/wCM+v8Ai/r/AIv6/zWk+v9RsPn/AIv7/wCM+v8AjPr/AIv6/wCM+v8AjPr/AIv6/weP+f9Cqfn/AIv6/wCL+v8Ai/r/AIz7/wCM+v8AjPv/AIv6/zak+P9RsPn/AIv6/wCM+/8AjPv/AIv6/wCM+v8Kj/n/kM75/y+g+P8Ai/r/AIv7/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIv6/wCM+/8Ai/r/AIv7/wCL+/8UlPn/mNH5/x2Y+f8Ai/v/AIv6/wCM+/8AjPv/AIv6/zWk+v9RsPn/AIv7/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/yed+f9uvfr/AIz6/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/zak+P9RsPn/AIz7/wCM+/8Ai/r/AIv6/xGS+f+X0Pn/I5r5/wCL+v8Ai/v/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8Ai/v/DJH5/5TP+f8pnfj/AIv6/wCM+/8AjPv/AIv6/zWk+v9RsPn/AIv7/wCM+/8AjPv/AIz7/wCM+/8Ai/v/AIv7/1Kw+f+j1fr/AIv6/wCM+/8AjPr/AIz7/wCM+v8AjPv/AIz6/zak+P9RsPn/AIz6/wCL+/8Ai/v/G5b5/5jQ+f8Xlfn/AIv6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz6/wiO+f+Ly/n/N6P5/wCL+v8Ai/v/AIv6/zWk+v9RsPn/AIv7/wCM+/8AjPv/AIz6/wCM+/8AjPr/AIv6/4nJ+v/i8vz/A4z5/wCL+v8Ai/r/AIz7/wCM+v8Ai/v/AIz6/zak+P9RsPn/AIz6/wGM+v8lnPj/ldD5/xGS+P8Ai/v/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8DjPr/gMb4/0ar+P8Ai/r/AIv6/zWk+v9RsPn/AIv7/wCM+/8AjPv/AIz6/wCM+/8Ai/r/AYv6/8rn/P/9/v3/NqP5/wCL+v8Ai/r/AIv7/wCL+v8Ai/v/AIv6/zak+P9SsPn/AYv6/zKh+f+Qzfj/CY75/wCM+/8Ai/v/AIv6/wCL+v8Ai/r/AIv6/wCL+/8Ai/v/AIv6/wCL+v8Ai/v/AIv6/wCL+v8Ai/v/AIv6/wCL+/8Ai/v/AIv6/wCL+v8AjPr/Aov6/3K/+f9Wsvj/AYz6/zWk+v9RsPn/AIv7/wCL+/8Ai/v/AIv6/wCL+/8Ai/r/Hpj5//r9/v/9/v7/kM36/wGL+/8AjPr/AIz7/wCM+v8AjPr/AIz6/zak+P9SsPj/Qaj4/4XH+P8Fjfn/AIv6/wCL+/8AjPv/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPv/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIv7/wGL+v9ht/j/aLr4/zak+f9RsPn/AIv6/wCM+v8AjPv/AIz6/wCL+/8AjPr/dMD5//7+/v//////6/b9/w2R+f8AjPr/AIz6/wCM+v8AjPv/AIz7/zKi+P+f1Pf/d8H4/wKM+f8Bi/r/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCM+/8AjPr/UK/4/6bW+P9Nrvn/AIv6/wCL+/8AjPv/AIz6/wCL+v8Fjfn/1+37//3+/v///////v7+/2+9+f8Bi/r/AIv7/wCL+v8Ai/v/AIv6/wWN+f+Oy/j/iMn5/4fJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/4jJ+f+Iyfn/iMn5/5HM+f8Nkfn/AIv7/wCL+/8Ai/v/AIv6/wCL+v9Tsfn//v7+//7+/v///////v7+/+Tz/P8Pkfn/AIv7/wCM+/8AjPv/AIz7/wCL+/8Ai/r/AIv7/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIv6/wCL+v8Ai/v/AIv6/wCM+/8AjPv/AIv7/waN+f/R6vv//v7+//////////////////7+/v+HyPn/AIv7/wCL+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/2y8+f/8/v3//v7+///////+/v7//v7+//7+/v/3/P3/NKP4/wCL+v8AjPr/AIv6/wCL+v8AjPv/AIv6/wCM+v8AjPr/AIv6/wCM+/8Ai/r/AIv6/wCM+v8AjPr/AIv6/wCL+v8AjPr/AIz6/wCL+v8Ai/r/AIv6/wCL+v8AjPr/AIv6/wCM+v8AjPr/AIv6/wCL+v8AjPr/AIv6/wCM+/8AjPr/AIv6/wCM+/8Ai/r/AIv6/wCM+v8Ai/r/Ipr4/+74/f/+/v7//v7+//7+/v/////////////+/v/+/v7/1ez7/xCR+f8Bi/r/AIv7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8Ai/v/AIv6/wCM+/8Ai/v/AIv7/wCM+/8Ai/v/AIv7/wCM+/8AjPv/AIz7/wCL+/8Hjvn/wuP7//3+/v///////v7+/////////////////////////////v79/7Db+f8Gjfr/AIv6/wCL+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz6/wCL+/8AjPv/AIz7/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wGM+v+Xz/n//f7+//7+/v///////v7+///////+/v7//v7+//7+/v/+/v7//v7+//7+/v+Y0Pn/A4z6/wCL+v8AjPr/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+v8AjPv/AIz6/wCM+v8AjPv/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIz6/wCM+v8AjPr/AIv6/wCM+v8AjPv/AIv6/wCM+v8AjPr/AIv6/wCM+/8Ai/v/AYv6/37E+f/8/v7//v7+//7+/v/+/v7//v7+//7+/v////////////////////////////7+/v/+/v7/mND5/waN+v8Bi/r/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8CjPr/f8X4//z+/f/+/v7//v/+/////////////v7+///////+/v7////////////+/v7///////7+/v/+/v7//v7+/7Db+f8Qkvn/AIv6/wCL+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCL+/8AjPv/AIz7/wCL+/8AjPv/AIz7/wCL+/8AjPv/AIz7/wCL+/8AjPv/AIz7/wCL+/8AjPv/AIz6/wiP+f+a0fn//f79//7+/v///////v7+/////////////v7+///////+/v7///////7+/v/+/v7//v7+//7+/v///////v7+//7+/v/V7Pv/NaP4/wCL+/8Ai/v/AIv7/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIv6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCM+/8AjPv/AIz6/wCL+/8Ai/v/J534/8jm+//9/v7//v7+//7+/v/+/v7//v7+//7+/v/+/v7//v7+///////+/v7////////////+/v7///////7+/v///////v7+///////+/v7/9/v9/4fI+f8Pkfn/AYv6/wCL+/8Ai/r/AIv6/wCL+v8Ai/r/AIv6/wCL+v8Ai/r/AIv6/wCL+v8Ai/r/AIv6/wCL+v8Ai/r/AIv6/wCL+v8Ai/r/AIv6/wCL+v8Ai/r/AIv7/wiP+f92wfn/8vn9/////////////v7+/////////////v7+///////+/v7//v7+///////+/v7////////////+/v7///////7+/v///////v7+/////////v7//v7+//7+/v/k8/3/br35/w2R+f8Bi/r/AIv7/wCL+/8Ai/v/AIz6/wCL+/8Ai/v/AIz6/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wCM+/8AjPv/AIz7/wGL+v8Jjvn/Ybf5/9ju+//9/v3//v7+/////////////v7+/////////////v7+/////////////v7+/////////////////////////////////////////////v7+/////////////v7+///////+/v7//f7+/+z2/f+QzPr/N6T5/wOM+f8Ai/v/AIz6/wCL+/8AjPr/AIz6/wCL+/8Ai/r/AIz6/wCL+/8Ai/r/AIz6/wCL+/8CjPr/LaD4/4bI+v/j8/z//v79//7+/v/+/v7//v7+/////////////////////////////v7+/////////////v7+///////+/v7////////////+/v7//////////////////v7+/////////////v7+/////////////v7+///////+/v7//P79/+Ly/P+k1fr/brz5/0Op+f8jm/j/DpH4/wKN+v8BjPv/DJD5/yCa+f8/qPn/aLr5/53S+v/b7/v/+/79//7+/v///////v7+/////////////v7+/////////////v7+/////////////v7+/////////////v7+//////8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
	CH         = make(chan bool, 100)
	stateurl   = "http://localhost/message"
	dataurl    = "http://localhost"
)

func Message(c unsafe.Pointer, text, caption string) {
	Text := syscall.StringToUTF16Ptr(text)
	Caption := syscall.StringToUTF16Ptr(caption)
	handle := win.HWND(uintptr(c))
	ret := win.MessageBox(handle, Text, Caption, win.MB_YESNO)
	if ret == win.IDYES {
		fmt.Println("点击YES")
	} else if ret == win.IDNO {
		fmt.Println("点击NO")
	}
}
func ShowWindow(c unsafe.Pointer) {
	handle := win.HWND(uintptr(c))
	win.ShowWindow(handle, win.SW_SHOW)
}

func HideWindow(c unsafe.Pointer) {
	handle := win.HWND(uintptr(c))
	win.ShowWindow(handle, win.SW_HIDE)
}
func MoveToCenter(c unsafe.Pointer) {
	handle := win.HWND(uintptr(c))
	var width int32 = 0
	var height int32 = 0
	{
		rect := &win.RECT{}
		win.GetWindowRect(handle, rect)
		width = rect.Right - rect.Left
		height = rect.Bottom - rect.Top
	}

	var parentWidth int32 = 0
	var parentHeight int32 = 0
	if win.GetWindowLong(handle, win.GWL_STYLE) == win.WS_CHILD {
		parent := win.GetParent(handle)
		rect := &win.RECT{}
		win.GetClientRect(parent, rect)
		parentWidth = rect.Right - rect.Left
		parentHeight = rect.Bottom - rect.Top
	} else {
		parentWidth = win.GetSystemMetrics(win.SM_CXSCREEN)
		parentHeight = win.GetSystemMetrics(win.SM_CYSCREEN)
	}

	x := (parentWidth - width) / 2
	y := (parentHeight - height) / 2
	//handle := win.HWND(uintptr(unsafe.Pointer(C.getWindowHandle(view.window))))
	win.MoveWindow(handle, x, y, width, height, false)
}
func Show() {
	go func() {
		w := webview.New(false)
		defer w.Destroy()
		w.SetTitle("data center")
		w.SetSize(1062, 960, webview.HintNone)
		//w.SetHtml("Thanks for using webview!")
		w.Navigate(dataurl)
		c := w.Window()
		MoveToCenter(c)
		//Message(c, "test", "test")
		w.Run()
	}()
}
func onReady() {
	ico, err := base64.StdEncoding.DecodeString(icobase64)
	if err == nil {
		systray.SetIcon(ico)
	}
	go func() {
		bl := true
		for {
			select {
			case <-CH:
				if bl {
					ico2, err := base64.StdEncoding.DecodeString(ico2base64)
					if err == nil {
						systray.SetIcon(ico2)
					}
					bl = false
				} else {
					ico, err := base64.StdEncoding.DecodeString(icobase64)
					if err == nil {
						systray.SetIcon(ico)
					}
					bl = true
				}
				CH <- true
				time.Sleep(500 * time.Millisecond)

			}
		}
	}()
	submenu := systray.AddMenuItem("显示", "显示主页", 0)

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出应用", 0)

	go func() {
		for {
			select {
			case <-submenu.OnClickCh():
				Show()
				clearCH()
			case <-mQuit.OnClickCh():
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Cleaning stuff here.
}
func getData() {
	go func() {
		for range time.Tick(5 * time.Second) {
			res, err := http.Get(stateurl)
			if err == nil {
				bs, err := ioutil.ReadAll(res.Body)
				if err == nil {
					rs := string(bs)
					if rs == "true" {
						CH <- true
					}
				}
				res.Body.Close()
			}
		}
	}()
}
func clearCH() {
	for i := 0; i < len(CH); i++ {
		<-CH
	}
}
func main() {
	getData()
	Show()
	systray.Run(onReady, onExit)
}
