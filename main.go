package main

import (
	"encoding/base64"

	"fmt"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"github.com/riftbit/go-systray"
	"github.com/webview/webview"
)

var (
	icobase64 = "AAABAAEALi4AAAEAIACoIgAAFgAAACgAAAAuAAAAXAAAAAEAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADCvLj/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/CvLj/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////NP///zT///80////NP///zT///81/Pz/N/Dv/z3Kxf9LfXL/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP39/zzRzP9PZVf/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///80////OeXi/09qXf9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////OOvp/zvV0f871dH/O9XR/znl4v81/Pz/NP///zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///z/Cvf9WQC//VkAv/1ZAL/9WQC//Tm5i/zre2/80////NP///zT///80////N+/u/0x2av9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VkAv/1ZAL/9KgXb/NvTz/zT///80////NP///zX6+f9HlYv/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////P8K9/1ZAL/9WQC//VkAv/1ZAL/9WQC//T2ZZ/zjm5P80////NP///zT///80////RKeg/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///z/Cvf9WQC//VkAv/1ZAL/9WQC//VkAv/09oWv845+X/NP///zT///80////NP39/0WhmP9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VkAv/1ZAL/9DrKT/Nfz8/zT///80////NP///zb08/9KgXb/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////PsnE/0x4bP9MeGz/THhs/0iMgv880cz/Nfv6/zT///80////NP///zT///8+ycT/VEk5/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///863dn/T2pd/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zfx7/9Drab/U1BB/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////NP///zT///80////NP///zT///80////NP39/zva1v9KgXb/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///zva1v9DraX/Qq+o/z3NyP837Or/NP7+/zT///80////Nfz8/0C8tv9USTn/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VEk5/0G2sP80/f3/NP///zT///80/v7/P8C6/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////P8K9/1ZAL/9WQC//VkAv/1ZAL/9PZln/OObk/zT///80////NP///zb08/9KgXb/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///z/Cvf9WQC//VkAv/1ZAL/9WQC//VkAv/zvY1f80////NP///zT///80/v7/RKWd/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///8/wr3/VkAv/1ZAL/9WQC//VkAv/01wY/846uj/NP///zT///80////NP///0SnoP9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////P8K9/1ZAL/9WQC//VkAv/0t9cv872tf/NP7+/zT///80////NP///zT9/f9FoZj/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9Ep6D/NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///837uz/TXRo/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//RKeg/zT///80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///81+/r/Q6uk/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/0SnoP80////NP///zT///80////NP///zT///80////NP///zT///80////NP///zT///846uj/SYqA/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9IkIb/O9XR/zvV0f871dH/O9XR/zvV0f871dH/O9XR/zvV0f871dH/PNPP/0G0rf9LgHX/VEk5/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/m5ub/VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL//m5ub/5ubm/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//5ubm/+bm5v9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/1ZAL/9WQC//VkAv/+bm5v/CvLj/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/m5ub/5ubm/+bm5v/CvLj/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA="
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
		w.Navigate("http://www.baidu.com")
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
	submenu := systray.AddMenuItem("显示", "显示主页", 0)

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出应用", 0)

	go func() {
		for {
			select {
			case <-submenu.OnClickCh():
				Show()
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
func main() {
	Show()
	systray.Run(onReady, onExit)
}
