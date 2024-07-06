// +build !noasm,!gccgo,!safe

#include "textflag.h"

#define X_PTR SI
#define IDX AX
#define LEN CX
#define TAIL BX
#define SUM X0
#define SUM_1 X1
#define SUM_2 X2
#define SUM_3 X3

// func Sum32(x []float32) float32
TEXT ·Sum32(SB), NOSPLIT, $0
    MOVQ x_base+0(FP), X_PTR // X_PTR = &x
    MOVQ x_len+8(FP), LEN // LEN = len(x)
    XORQ IDX, IDX // i = 0
    PXOR SUM, SUM // p_sum_i = 0
    CMPQ LEN, $0 // if LEN == 0 { return 0 }
    JE sum_end

    PXOR SUM_1, SUM_1
    PXOR SUM_2, SUM_2
    PXOR SUM_3, SUM_3

    MOVQ X_PTR, TAIL // Check memory alignment
    ANDQ $15, TAIL // TAIL = &x % 16
    JZ no_trim // if TAIL == 0 { goto no_trim }
    SUBQ $16, TAIL // TAIL -= 16

sum_align: // Align on 16-byte boundary do {
    ADDSS (X_PTR)(IDX*4), SUM // SUM += x[0]
    INCQ IDX // i++
    DECQ LEN // LEN--
    JZ sum_end // if LEN == 0 { return }
    ADDQ $4, TAIL // TAIL += 4
    JNZ sum_align // } while TAIL < 0

no_trim:
    MOVQ LEN, TAIL
    SHRQ $4, LEN // LEN = floor( n / 16 )
    JZ sum_tail8 // if LEN == 0 { goto sum_tail8 }

sum_loop: // sum 16x wide do {
    ADDPS (X_PTR)(IDX*4), SUM // sum_i += x[i:i+4]
    ADDPS 16(X_PTR)(IDX*4), SUM_1
    ADDPS 32(X_PTR)(IDX*4), SUM_2
    ADDPS 48(X_PTR)(IDX*4), SUM_3

    ADDQ $16, IDX // i += 16
    DECQ LEN
    JNZ sum_loop // } while —LEN > 0

sum_tail8:
    ADDPS SUM_3, SUM
    ADDPS SUM_2, SUM_1

    TESTQ $8, TAIL
    JZ sum_tail4

    ADDPS (X_PTR)(IDX*4), SUM // sum_i += x[i:i+4]
    ADDPS 16(X_PTR)(IDX*4), SUM_1
    ADDQ $8, IDX

sum_tail4:
    ADDPS SUM_1, SUM

    TESTQ $4, TAIL
    JZ sum_tail2

    ADDPS (X_PTR)(IDX*4), SUM // sum_i += x[i:i+4]
    ADDQ $4, IDX

sum_tail2:
    HADDPS SUM, SUM // sum_i[:2] += sum_i[2:4]

    TESTQ $2, TAIL
    JZ sum_tail1

    MOVSD (X_PTR)(IDX*4), SUM_1 // reuse SUM_1
    ADDPS SUM_1, SUM // sum_i += x[i:i+2]
    ADDQ $2, IDX

sum_tail1:
    HADDPS SUM, SUM // sum_i[0] += sum_i[1]

    TESTQ $1, TAIL
    JZ sum_end

    ADDSS (X_PTR)(IDX*4), SUM

    sum_end: // return sum
    MOVSS SUM, ret+24(FP)
    RET

// func Mul32(x []float32, y []float32)
// Requires: AVX
TEXT ·Mul32(SB), NOSPLIT, $0-48
    MOVQ x_base+0(FP), DI
    MOVQ y_base+24(FP), SI
    MOVQ x_len+8(FP), DX
    TESTQ DX, DX
    JE LBB9_7
    CMPQ DX, $0x20
    JAE LBB9_3
    XORL AX, AX
    JMP LBB9_6

LBB9_3:
    MOVQ DX, AX
    ANDQ $-32, AX
    XORL CX, CX

LBB9_4:
    VMOVUPS (DI)(CX*4), Y0
    VMOVUPS 32(DI)(CX*4), Y1
    VMOVUPS 64(DI)(CX*4), Y2
    VMOVUPS 96(DI)(CX*4), Y3
    VMULPS (SI)(CX*4), Y0, Y0
    VMULPS 32(SI)(CX*4), Y1, Y1
    VMULPS 64(SI)(CX*4), Y2, Y2
    VMULPS 96(SI)(CX*4), Y3, Y3
    VMOVUPS Y0, (DI)(CX*4)
    VMOVUPS Y1, 32(DI)(CX*4)
    VMOVUPS Y2, 64(DI)(CX*4)
    VMOVUPS Y3, 96(DI)(CX*4)
    ADDQ $0x20, CX
    CMPQ AX, CX
    JNE LBB9_4
    CMPQ AX, DX
    JE LBB9_7

LBB9_6:
    VMOVSS (DI)(AX*4), X0
    VMULSS (SI)(AX*4), X0, X0
    VMOVSS X0, (DI)(AX*4)
    ADDQ $0x01, AX
    CMPQ DX, AX
    JNE LBB9_6

LBB9_7:
    VZEROUPPER
    RET

// func Mul64(x []float64, y []float64)
// Requires: AVX
TEXT ·Mul64(SB), NOSPLIT, $0-48
    MOVQ x_base+0(FP), DI
    MOVQ y_base+24(FP), SI
    MOVQ x_len+8(FP), DX
    TESTQ DX, DX
    JE LBB8_7
    CMPQ DX, $0x10
    JAE LBB8_3
    XORL AX, AX
    JMP LBB8_6

    LBB8_3:
    MOVQ DX, AX
    ANDQ $-16, AX
    XORL CX, CX

LBB8_4:
    VMOVUPD (DI)(CX*8), Y0
    VMOVUPD 32(DI)(CX*8), Y1
    VMOVUPD 64(DI)(CX*8), Y2
    VMOVUPD 96(DI)(CX*8), Y3
    VMULPD (SI)(CX*8), Y0, Y0
    VMULPD 32(SI)(CX*8), Y1, Y1
    VMULPD 64(SI)(CX*8), Y2, Y2
    VMULPD 96(SI)(CX*8), Y3, Y3
    VMOVUPD Y0, (DI)(CX*8)
    VMOVUPD Y1, 32(DI)(CX*8)
    VMOVUPD Y2, 64(DI)(CX*8)
    VMOVUPD Y3, 96(DI)(CX*8)
    ADDQ $0x10, CX
    CMPQ AX, CX
    JNE LBB8_4
    CMPQ AX, DX
    JE LBB8_7

LBB8_6:
    VMOVSD (DI)(AX*8), X0
    VMULSD (SI)(AX*8), X0, X0
    VMOVSD X0, (DI)(AX*8)
    ADDQ $0x01, AX
    CMPQ DX, AX
    JNE LBB8_6

LBB8_7:
    VZEROUPPER
    RET

// func MulNum32(x []float32, a float32)
// Requires: AVX, AVX2, SSE
TEXT ·MulNum32(SB), NOSPLIT, $0-28
    MOVQ x_base+0(FP), DI
    MOVSS a+24(FP), X0
    MOVQ x_len+8(FP), SI
    TESTQ SI, SI
    JE LBB11_11
    CMPQ SI, $0x20
    JAE LBB11_3
    XORL AX, AX
    JMP LBB11_10

LBB11_3:
    MOVQ SI, AX
    ANDQ $-32, AX
    VBROADCASTSS X0, Y1
    LEAQ -32(AX), CX
    MOVQ CX, R8
    SHRQ $0x05, R8
    ADDQ $0x01, R8
    TESTQ CX, CX
    JE LBB11_4
    MOVQ R8, DX
    ANDQ $-2, DX
    XORL CX, CX

LBB11_6:
    VMULPS (DI)(CX*4), Y1, Y2
    VMULPS 32(DI)(CX*4), Y1, Y3
    VMULPS 64(DI)(CX*4), Y1, Y4
    VMULPS 96(DI)(CX*4), Y1, Y5
    VMOVUPS Y2, (DI)(CX*4)
    VMOVUPS Y3, 32(DI)(CX*4)
    VMOVUPS Y4, 64(DI)(CX*4)
    VMOVUPS Y5, 96(DI)(CX*4)
    VMULPS 128(DI)(CX*4), Y1, Y2
    VMULPS 160(DI)(CX*4), Y1, Y3
    VMULPS 192(DI)(CX*4), Y1, Y4
    VMULPS 224(DI)(CX*4), Y1, Y5
    VMOVUPS Y2, 128(DI)(CX*4)
    VMOVUPS Y3, 160(DI)(CX*4)
    VMOVUPS Y4, 192(DI)(CX*4)
    VMOVUPS Y5, 224(DI)(CX*4)
    ADDQ $0x40, CX
    ADDQ $-2, DX
    JNE LBB11_6
    TESTB $0x01, R8
    JE LBB11_9

LBB11_8:
    VMULPS (DI)(CX*4), Y1, Y2
    VMULPS 32(DI)(CX*4), Y1, Y3
    VMULPS 64(DI)(CX*4), Y1, Y4
    VMULPS 96(DI)(CX*4), Y1, Y1
    VMOVUPS Y2, (DI)(CX*4)
    VMOVUPS Y3, 32(DI)(CX*4)
    VMOVUPS Y4, 64(DI)(CX*4)
    VMOVUPS Y1, 96(DI)(CX*4)

LBB11_9:
    CMPQ AX, SI
    JE LBB11_11

LBB11_10:
    VMULSS (DI)(AX*4), X0, X1
    VMOVSS X1, (DI)(AX*4)
    ADDQ $0x01, AX
    CMPQ SI, AX
    JNE LBB11_10

LBB11_11:
    VZEROUPPER
    RET

LBB11_4:
    XORL CX, CX
    TESTB $0x01, R8
    JNE LBB11_8
    JMP LBB11_9

// func MulNum64(x []float64, a float64)
// Requires: AVX, AVX2, SSE2
TEXT ·MulNum64(SB), NOSPLIT, $0-32
    MOVQ x_base+0(FP), DI
    MOVSD a+24(FP), X0
    MOVQ x_len+8(FP), SI
    TESTQ SI, SI
    JE LBB10_11
    CMPQ SI, $0x10
    JAE LBB10_3
    XORL AX, AX
    JMP LBB10_10

LBB10_3:
    MOVQ SI, AX
    ANDQ $-16, AX
    VBROADCASTSD X0, Y1
    LEAQ -16(AX), CX
    MOVQ CX, R8
    SHRQ $0x04, R8
    ADDQ $0x01, R8
    TESTQ CX, CX
    JE LBB10_4
    MOVQ R8, DX
    ANDQ $-2, DX
    XORL CX, CX

LBB10_6:
    VMULPD (DI)(CX*8), Y1, Y2
    VMULPD 32(DI)(CX*8), Y1, Y3
    VMULPD 64(DI)(CX*8), Y1, Y4
    VMULPD 96(DI)(CX*8), Y1, Y5
    VMOVUPD Y2, (DI)(CX*8)
    VMOVUPD Y3, 32(DI)(CX*8)
    VMOVUPD Y4, 64(DI)(CX*8)
    VMOVUPD Y5, 96(DI)(CX*8)
    VMULPD 128(DI)(CX*8), Y1, Y2
    VMULPD 160(DI)(CX*8), Y1, Y3
    VMULPD 192(DI)(CX*8), Y1, Y4
    VMULPD 224(DI)(CX*8), Y1, Y5
    VMOVUPD Y2, 128(DI)(CX*8)
    VMOVUPD Y3, 160(DI)(CX*8)
    VMOVUPD Y4, 192(DI)(CX*8)
    VMOVUPD Y5, 224(DI)(CX*8)
    ADDQ $0x20, CX
    ADDQ $-2, DX
    JNE LBB10_6
    TESTB $0x01, R8
    JE LBB10_9

LBB10_8:
    VMULPD (DI)(CX*8), Y1, Y2
    VMULPD 32(DI)(CX*8), Y1, Y3
    VMULPD 64(DI)(CX*8), Y1, Y4
    VMULPD 96(DI)(CX*8), Y1, Y1
    VMOVUPD Y2, (DI)(CX*8)
    VMOVUPD Y3, 32(DI)(CX*8)
    VMOVUPD Y4, 64(DI)(CX*8)
    VMOVUPD Y1, 96(DI)(CX*8)

LBB10_9:
    CMPQ AX, SI
    JE LBB10_11

LBB10_10:
    VMULSD (DI)(AX*8), X0, X1
    VMOVSD X1, (DI)(AX*8)
    ADDQ $0x01, AX
    CMPQ SI, AX
    JNE LBB10_10

LBB10_11:
    VZEROUPPER
    RET

LBB10_4:
    XORL CX, CX
    TESTB $0x01, R8
    JNE LBB10_8
    JMP LBB10_9

// func Add32(x []float32, y []float32)
// Requires: AVX
TEXT ·Add32(SB), NOSPLIT, $0-48
    MOVQ x_base+0(FP), DI
    MOVQ y_base+24(FP), SI
    MOVQ x_len+8(FP), DX
    TESTQ DX, DX
    JE LBB1_7
    CMPQ DX, $0x20
    JAE LBB1_3
    XORL AX, AX
    JMP LBB1_6

LBB1_3:
    MOVQ DX, AX
    ANDQ $-32, AX
    XORL CX, CX

LBB1_4:
    VMOVUPS (DI)(CX*4), Y0
    VMOVUPS 32(DI)(CX*4), Y1
    VMOVUPS 64(DI)(CX*4), Y2
    VMOVUPS 96(DI)(CX*4), Y3
    VADDPS (SI)(CX*4), Y0, Y0
    VADDPS 32(SI)(CX*4), Y1, Y1
    VADDPS 64(SI)(CX*4), Y2, Y2
    VADDPS 96(SI)(CX*4), Y3, Y3
    VMOVUPS Y0, (DI)(CX*4)
    VMOVUPS Y1, 32(DI)(CX*4)
    VMOVUPS Y2, 64(DI)(CX*4)
    VMOVUPS Y3, 96(DI)(CX*4)
    ADDQ $0x20, CX
    CMPQ AX, CX
    JNE LBB1_4
    CMPQ AX, DX
    JE LBB1_7

LBB1_6:
    VMOVSS (DI)(AX*4), X0
    VADDSS (SI)(AX*4), X0, X0
    VMOVSS X0, (DI)(AX*4)
    ADDQ $0x01, AX
    CMPQ DX, AX
    JNE LBB1_6

LBB1_7:
    VZEROUPPER
    RET

// func Add64(x []float64, y []float64)
// Requires: AVX
TEXT ·Add64(SB), NOSPLIT, $0-48
    MOVQ x_base+0(FP), DI
    MOVQ y_base+24(FP), SI
    MOVQ x_len+8(FP), DX
    TESTQ DX, DX
    JE LBB0_7
    CMPQ DX, $0x10
    JAE LBB0_3
    XORL AX, AX
    JMP LBB0_6

LBB0_3:
    MOVQ DX, AX
    ANDQ $-16, AX
    XORL CX, CX

LBB0_4:
    VMOVUPD (DI)(CX*8), Y0
    VMOVUPD 32(DI)(CX*8), Y1
    VMOVUPD 64(DI)(CX*8), Y2
    VMOVUPD 96(DI)(CX*8), Y3
    VADDPD (SI)(CX*8), Y0, Y0
    VADDPD 32(SI)(CX*8), Y1, Y1
    VADDPD 64(SI)(CX*8), Y2, Y2
    VADDPD 96(SI)(CX*8), Y3, Y3
    VMOVUPD Y0, (DI)(CX*8)
    VMOVUPD Y1, 32(DI)(CX*8)
    VMOVUPD Y2, 64(DI)(CX*8)
    VMOVUPD Y3, 96(DI)(CX*8)
    ADDQ $0x10, CX
    CMPQ AX, CX
    JNE LBB0_4
    CMPQ AX, DX
    JE LBB0_7

LBB0_6:
    VMOVSD (DI)(AX*8), X0
    VADDSD (SI)(AX*8), X0, X0
    VMOVSD X0, (DI)(AX*8)
    ADDQ $0x01, AX
    CMPQ DX, AX
    JNE LBB0_6

LBB0_7:
    VZEROUPPER
    RET

// func AddNum32(x []float32, a float32)
// Requires: AVX, AVX2, SSE
TEXT ·AddNum32(SB), NOSPLIT, $0-28
    MOVQ x_base+0(FP), DI
    MOVSS a+24(FP), X0
    MOVQ x_len+8(FP), SI
    TESTQ SI, SI
    JE LBB3_11
    CMPQ SI, $0x20
    JAE LBB3_3
    XORL AX, AX
    JMP LBB3_10

LBB3_3:
    MOVQ SI, AX
    ANDQ $-32, AX
    VBROADCASTSS X0, Y1
    LEAQ -32(AX), CX
    MOVQ CX, R8
    SHRQ $0x05, R8
    ADDQ $0x01, R8
    TESTQ CX, CX
    JE LBB3_4
    MOVQ R8, DX
    ANDQ $-2, DX
    XORL CX, CX

LBB3_6:
    VADDPS (DI)(CX*4), Y1, Y2
    VADDPS 32(DI)(CX*4), Y1, Y3
    VADDPS 64(DI)(CX*4), Y1, Y4
    VADDPS 96(DI)(CX*4), Y1, Y5
    VMOVUPS Y2, (DI)(CX*4)
    VMOVUPS Y3, 32(DI)(CX*4)
    VMOVUPS Y4, 64(DI)(CX*4)
    VMOVUPS Y5, 96(DI)(CX*4)
    VADDPS 128(DI)(CX*4), Y1, Y2
    VADDPS 160(DI)(CX*4), Y1, Y3
    VADDPS 192(DI)(CX*4), Y1, Y4
    VADDPS 224(DI)(CX*4), Y1, Y5
    VMOVUPS Y2, 128(DI)(CX*4)
    VMOVUPS Y3, 160(DI)(CX*4)
    VMOVUPS Y4, 192(DI)(CX*4)
    VMOVUPS Y5, 224(DI)(CX*4)
    ADDQ $0x40, CX
    ADDQ $-2, DX
    JNE LBB3_6
    TESTB $0x01, R8
    JE LBB3_9

LBB3_8:
    VADDPS (DI)(CX*4), Y1, Y2
    VADDPS 32(DI)(CX*4), Y1, Y3
    VADDPS 64(DI)(CX*4), Y1, Y4
    VADDPS 96(DI)(CX*4), Y1, Y1
    VMOVUPS Y2, (DI)(CX*4)
    VMOVUPS Y3, 32(DI)(CX*4)
    VMOVUPS Y4, 64(DI)(CX*4)
    VMOVUPS Y1, 96(DI)(CX*4)

LBB3_9:
    CMPQ AX, SI
    JE LBB3_11

LBB3_10:
    VADDSS (DI)(AX*4), X0, X1
    VMOVSS X1, (DI)(AX*4)
    ADDQ $0x01, AX
    CMPQ SI, AX
    JNE LBB3_10

LBB3_11:
    VZEROUPPER
    RET

LBB3_4:
    XORL CX, CX
    TESTB $0x01, R8
    JNE LBB3_8
    JMP LBB3_9

// func AddNum64(x []float64, a float64)
// Requires: AVX, AVX2, SSE2
TEXT ·AddNum64(SB), NOSPLIT, $0-32
    MOVQ x_base+0(FP), DI
    MOVSD a+24(FP), X0
    MOVQ x_len+8(FP), SI
    TESTQ SI, SI
    JE LBB2_11
    CMPQ SI, $0x10
    JAE LBB2_3
    XORL AX, AX
    JMP LBB2_10

LBB2_3:
    MOVQ SI, AX
    ANDQ $-16, AX
    VBROADCASTSD X0, Y1
    LEAQ -16(AX), CX
    MOVQ CX, R8
    SHRQ $0x04, R8
    ADDQ $0x01, R8
    TESTQ CX, CX
    JE LBB2_4
    MOVQ R8, DX
    ANDQ $-2, DX
    XORL CX, CX

LBB2_6:
    VADDPD (DI)(CX*8), Y1, Y2
    VADDPD 32(DI)(CX*8), Y1, Y3
    VADDPD 64(DI)(CX*8), Y1, Y4
    VADDPD 96(DI)(CX*8), Y1, Y5
    VMOVUPD Y2, (DI)(CX*8)
    VMOVUPD Y3, 32(DI)(CX*8)
    VMOVUPD Y4, 64(DI)(CX*8)
    VMOVUPD Y5, 96(DI)(CX*8)
    VADDPD 128(DI)(CX*8), Y1, Y2
    VADDPD 160(DI)(CX*8), Y1, Y3
    VADDPD 192(DI)(CX*8), Y1, Y4
    VADDPD 224(DI)(CX*8), Y1, Y5
    VMOVUPD Y2, 128(DI)(CX*8)
    VMOVUPD Y3, 160(DI)(CX*8)
    VMOVUPD Y4, 192(DI)(CX*8)
    VMOVUPD Y5, 224(DI)(CX*8)
    ADDQ $0x20, CX
    ADDQ $-2, DX
    JNE LBB2_6
    TESTB $0x01, R8
    JE LBB2_9

LBB2_8:
    VADDPD (DI)(CX*8), Y1, Y2
    VADDPD 32(DI)(CX*8), Y1, Y3
    VADDPD 64(DI)(CX*8), Y1, Y4
    VADDPD 96(DI)(CX*8), Y1, Y1
    VMOVUPD Y2, (DI)(CX*8)
    VMOVUPD Y3, 32(DI)(CX*8)
    VMOVUPD Y4, 64(DI)(CX*8)
    VMOVUPD Y1, 96(DI)(CX*8)

LBB2_9:
    CMPQ AX, SI
    JE LBB2_11

LBB2_10:
    VADDSD (DI)(AX*8), X0, X1
    VMOVSD X1, (DI)(AX*8)
    ADDQ $0x01, AX
    CMPQ SI, AX
    JNE LBB2_10

LBB2_11:
    VZEROUPPER
    RET

LBB2_4:
    XORL CX, CX
    TESTB $0x01, R8
    JNE LBB2_8
    JMP LBB2_9

// func MaximumNum32(x []float32, a float32)
// Requires: AVX, AVX2, SSE
TEXT ·MaximumNum32(SB), NOSPLIT, $0-28
    MOVQ x_base+0(FP), DI
    MOVSS a+24(FP), X0
    MOVQ x_len+8(FP), SI
    TESTQ SI, SI
    JE return
    CMPQ SI, $0x20
    JAE loop
    XORL AX, AX
    JMP tailbody

loop:
    MOVQ SI, AX
    ANDQ $-32, AX
    VBROADCASTSS X0, Y1
    LEAQ 96(DI), CX
    XORL DX, DX

body:
    VMOVUPS -96(CX)(DX*4), Y2
    VMOVUPS -64(CX)(DX*4), Y3
    VMOVUPS -32(CX)(DX*4), Y4
    VMOVUPS (CX)(DX*4), Y5
    VCMPPS $0x01, Y1, Y2, Y2
    VMASKMOVPS Y1, Y2, -96(CX)(DX*4)
    VCMPPS $0x01, Y1, Y3, Y2
    VMASKMOVPS Y1, Y2, -64(CX)(DX*4)
    VCMPPS $0x01, Y1, Y4, Y2
    VMASKMOVPS Y1, Y2, -32(CX)(DX*4)
    VCMPPS $0x01, Y1, Y5, Y2
    VMASKMOVPS Y1, Y2, (CX)(DX*4)
    ADDQ $0x20, DX
    CMPQ AX, DX
    JNE body
    CMPQ AX, SI
    JNE tailbody

return:
    VZEROUPPER
    RET

tail:
    ADDQ $0x01, AX
    CMPQ SI, AX
    JE return

tailbody:
    VUCOMISS (DI)(AX*4), X0
    JBE tail
    VMOVSS X0, (DI)(AX*4)
    JMP tail

// func MaximumNum64(x []float64, a float64)
// Requires: AVX, AVX2, SSE2
TEXT ·MaximumNum64(SB), NOSPLIT, $0-32
    MOVQ x_base+0(FP), DI
    MOVSD a+24(FP), X0
    MOVQ x_len+8(FP), SI
    TESTQ SI, SI
    JE return
    CMPQ SI, $0x10
    JAE loop
    XORL AX, AX
    JMP tailbody

loop:
    MOVQ SI, AX
    ANDQ $-16, AX
    VBROADCASTSD X0, Y1
    LEAQ 96(DI), CX
    XORL DX, DX

body:
    VMOVUPD -96(CX)(DX*8), Y2
    VMOVUPD -64(CX)(DX*8), Y3
    VMOVUPD -32(CX)(DX*8), Y4
    VMOVUPD (CX)(DX*8), Y5
    VCMPPD $0x01, Y1, Y2, Y2
    VMASKMOVPD Y1, Y2, -96(CX)(DX*8)
    VCMPPD $0x01, Y1, Y3, Y2
    VMASKMOVPD Y1, Y2, -64(CX)(DX*8)
    VCMPPD $0x01, Y1, Y4, Y2
    VMASKMOVPD Y1, Y2, -32(CX)(DX*8)
    VCMPPD $0x01, Y1, Y5, Y2
    VMASKMOVPD Y1, Y2, (CX)(DX*8)
    ADDQ $0x10, DX
    CMPQ AX, DX
    JNE body
    CMPQ AX, SI
    JNE tailbody

return:
    VZEROUPPER
    RET

tail:
    ADDQ $0x01, AX
    CMPQ SI, AX
    JE return

tailbody:
    VUCOMISD (DI)(AX*8), X0
    JBE tail
    VMOVSD X0, (DI)(AX*8)
    JMP tail

TEXT ·Sum64(SB), $0-24

    MOVQ buf+0(FP), DI
    MOVQ len+8(FP), SI
    MOVQ res+16(FP), DX

    WORD $0xf685 // test esi, esi
    JLE LBB0_1
    WORD $0x8941; BYTE $0xf1 // mov r9d, esi
    WORD $0xfe83; BYTE $0x0f // cmp esi, 15
    JA LBB0_4
    LONG $0xc057f9c5 // vxorpd xmm0, xmm0, xmm0
    WORD $0xc931 // xor ecx, ecx
    JMP LBB0_11

LBB0_1:
    LONG $0xc057f9c5 // vxorpd xmm0, xmm0, xmm0
    JMP LBB0_12

LBB0_4:
    WORD $0x8944; BYTE $0xc9 // mov ecx, r9d
    WORD $0xe183; BYTE $0xf0 // and ecx, -16
    LONG $0xf0718d48 // lea rsi, [rcx - 16]
    WORD $0x8948; BYTE $0xf0 // mov rax, rsi
    LONG $0x04e8c148 // shr rax, 4
    WORD $0xff48; BYTE $0xc0 // inc rax
    WORD $0x8941; BYTE $0xc0 // mov r8d, eax
    LONG $0x01e08341 // and r8d, 1
    WORD $0x8548; BYTE $0xf6 // test rsi, rsi
    JE LBB0_5
    LONG $0x000001be; BYTE $0x00 // mov esi, 1
    WORD $0x2948; BYTE $0xc6 // sub rsi, rax
    LONG $0x30448d49; BYTE $0xff // lea rax, [r8 + rsi - 1]
    LONG $0xc057f9c5 // vxorpd xmm0, xmm0, xmm0
    WORD $0xf631 // xor esi, esi
    LONG $0xc957f1c5 // vxorpd xmm1, xmm1, xmm1
    LONG $0xd257e9c5 // vxorpd xmm2, xmm2, xmm2
    LONG $0xdb57e1c5 // vxorpd xmm3, xmm3, xmm3

LBB0_7:
    LONG $0x0458fdc5; BYTE $0xf7 // vaddpd ymm0, ymm0, yword [rdi + 8*rsi]
    LONG $0x4c58f5c5; WORD $0x20f7 // vaddpd ymm1, ymm1, yword [rdi + 8*rsi + 32]
    LONG $0x5458edc5; WORD $0x40f7 // vaddpd ymm2, ymm2, yword [rdi + 8*rsi + 64]
    LONG $0x5c58e5c5; WORD $0x60f7 // vaddpd ymm3, ymm3, yword [rdi + 8*rsi + 96]
    QUAD $0x000080f78458fdc5; BYTE $0x00 // vaddpd ymm0, ymm0, yword [rdi + 8*rsi + 128]
    QUAD $0x0000a0f78c58f5c5; BYTE $0x00 // vaddpd ymm1, ymm1, yword [rdi + 8*rsi + 160]
    QUAD $0x0000c0f79458edc5; BYTE $0x00 // vaddpd ymm2, ymm2, yword [rdi + 8*rsi + 192]
    QUAD $0x0000e0f79c58e5c5; BYTE $0x00 // vaddpd ymm3, ymm3, yword [rdi + 8*rsi + 224]
    LONG $0x20c68348 // add rsi, 32
    LONG $0x02c08348 // add rax, 2
    JNE LBB0_7
    WORD $0x854d; BYTE $0xc0 // test r8, r8
    JE LBB0_10

LBB0_9:
    LONG $0x5c58e5c5; WORD $0x60f7 // vaddpd ymm3, ymm3, yword [rdi + 8*rsi + 96]
    LONG $0x5458edc5; WORD $0x40f7 // vaddpd ymm2, ymm2, yword [rdi + 8*rsi + 64]
    LONG $0x4c58f5c5; WORD $0x20f7 // vaddpd ymm1, ymm1, yword [rdi + 8*rsi + 32]
    LONG $0x0458fdc5; BYTE $0xf7 // vaddpd ymm0, ymm0, yword [rdi + 8*rsi]

LBB0_10:
    LONG $0xc058f5c5 // vaddpd ymm0, ymm1, ymm0
    LONG $0xc058edc5 // vaddpd ymm0, ymm2, ymm0
    LONG $0xc058e5c5 // vaddpd ymm0, ymm3, ymm0
    LONG $0x197de3c4; WORD $0x01c1 // vextractf128 xmm1, ymm0, 1
    LONG $0xc158fdc5 // vaddpd ymm0, ymm0, ymm1
    LONG $0xc07cfdc5 // vhaddpd ymm0, ymm0, ymm0
    WORD $0x394c; BYTE $0xc9 // cmp rcx, r9
    JE LBB0_12

LBB0_11:
    LONG $0x0458fbc5; BYTE $0xcf // vaddsd xmm0, xmm0, qword [rdi + 8*rcx]
    WORD $0xff48; BYTE $0xc1 // inc rcx
    WORD $0x3949; BYTE $0xc9 // cmp r9, rcx
    JNE LBB0_11

LBB0_12:
    LONG $0x0211fbc5 // vmovsd qword [rdx], xmm0
    VZEROUPPER
    RET

LBB0_5:
    LONG $0xc057f9c5 // vxorpd xmm0, xmm0, xmm0
    WORD $0xf631 // xor esi, esi
    LONG $0xc957f1c5 // vxorpd xmm1, xmm1, xmm1
    LONG $0xd257e9c5 // vxorpd xmm2, xmm2, xmm2
    LONG $0xdb57e1c5 // vxorpd xmm3, xmm3, xmm3
    WORD $0x854d; BYTE $0xc0 // test r8, r8
    JNE LBB0_9
    JMP LBB0_10

// func ·Mul64Simd(out []float64, x []float64, y []float64)
TEXT ·Mul64Simd(SB), 7, $0
    MOVQ out(FP),SI // SI: &out
    MOVQ out_len+8(FP),DX // DX: len(out)
    MOVQ x+24(FP),R11 // R11: &x
    MOVQ y+48(FP),R9 // R9: &y
    MOVQ DX, R10 // R10: len(out)
    SHRQ $2, DX // DX: len(out) / 4
    ANDQ $3, R10 // R10: len(out) % 4
    CMPQ DX ,$0
    JEQ remain_mul
loopback_mul:
    MOVUPD (R11),X0
    MOVUPD (R9),X1
    MULPD X1,X0
    MOVUPD 16(R11),X2
    MOVUPD 16(R9),X3
    MULPD X3,X2
    MOVUPD X0,(SI) //R11
    MOVUPD X2,16(SI) //R11
    ADDQ $32, R11
    ADDQ $32, R9
    ADDQ $32, SI
    SUBQ $1,DX
    JNZ loopback_mul
remain_mul:
    CMPQ R10,$0
    JEQ done_mul
onemore_mul:
    MOVSD (R11),X0
    MOVSD (R9),X1
    MULSD X1,X0
    MOVSD X0,(SI) //SI
    ADDQ $8, R11
    ADDQ $8, R9
    ADDQ $8, SI
    SUBQ $1, R10
    JNZ onemore_mul
done_mul:
    RET
