dict =
1. hariId # mengubah hari menjadi id
2. slotPerHari # mendefinisikan jumlah slot per hari
3. guruPengajar # mendefinikan nama pengajar
4. kelasDanTingkatan # mendefinisikan tingkatan kelas dan nama kelas
5. namaMapelDanId # mendefinisikan mapel id dan nama mapel
6. jamPerMingguMapel # mapel id dan jam per minggu
7. mgmpMapel # menyimpan mapel id dan hari MGMP
8. durasiGuruMengajar # guru id serta durasi mengajar
9. batasSiang # batas mapel id 8 diletakkan 
10. batasMGMP # mapel tidak boleh diletakkan di hari == MGMP, boleh tetapi tidak boleh melebihi slot ke 3

## 🔎 Daftar Constraint yang diSebutkan

### 1️⃣ Guru tidak boleh mengajar lebih dari 1 kelas pada slot waktu yang sama

➡ **Hard constraint**

---

### 2️⃣ Pola distribusi jam_per_minggu**

* jam=2 → [2] berurutan
* jam=3 → [3] berurutan
* jam=4 → [2,2] beda hari
* jam=5 → [3,2] beda hari

➡ Ini satu kelompok aturan distribusi
➡ Secara konsep ini dihitung sebagai **1 hard constraint (distribution rule)**

---

### 3️⃣ Mapel harus berdekatan slotnya pada hari yang sama

➡ Ini sebenarnya bagian dari constraint nomor 2
➡ Bukan constraint baru
➡ Jadi tidak dihitung terpisah

---

### 4️⃣ Mapel id=8 tidak boleh melewati batas siang

➡ **Hard constraint (time restriction)**

---

### 5️⃣ Durasi guru mengajar harus sesuai dictionary

➡ Jika artinya total jam guru tidak boleh melebihi batas →
➡ **Hard constraint (teacher load limit)**

---

### 6️⃣ jam_per_minggu tidak boleh melebihi dictionary

➡ Ini sebenarnya bagian dari constraint distribusi (no. 2)
➡ Jadi bukan constraint terpisah

---

### 7️⃣ MGMP constraint

* Tidak boleh hari tertentu
* Atau boleh tapi tidak lewat jam 09:00

➡ **Hard constraint (special scheduling rule)**

---

# 🎯 Jadi Total Hard Constraint Sebenarnya:

| No | Constraint                    | Hitung Terpisah? |
| -- | ----------------------------- | ---------------- |
| 1  | Guru bentrok                  | ✅                |
| 2  | Distribusi jam (blok & split) | ✅                |
| 3  | Batas siang mapel 8           | ✅                |
| 4  | Batas jam guru                | ✅                |
| 5  | MGMP rule                     | ✅                |

---

# ✅ Total = **5 Hard Constraint**

---
# representasi individu:
individu = {

'kelas_1': { # panjang array di tiap hari sesuai dengan jumlah slot per hari

    'Senin':  [1,1,2,2,2,8,8,8],     # 1 = block 2, 2 = block 3, 8 sebelum siang
    'Selasa': [4,4,4,3,3,5,5,5],     # 4 block 3, 3 block 2
    'Rabu':   [1,1,6,6,7,7,9,9],     # 1 split kedua (2 slot)
    'Kamis':  [6,6,10,10,11,11,12],  # 6 split kedua
    'Jumat':  [4,4,13,13,12]         # 4 split kedua (2 slot)
    }
}

# Penjelasan Distribusi Kelas_1
| Subject | Jam | Pola  | Terpenuhi?            |
| ------- | --- | ----- | --------------------- |
| 1       | 4   | [2,2] | Senin & Rabu          |
| 2       | 3   | [3]   | Senin                 |
| 3       | 2   | [2]   | Selasa                |
| 4       | 5   | [3,2] | Selasa & Jumat        |
| 6       | 4   | [2,2] | Rabu & Kamis          |
| 8       | 3   | [3]   | Senin (≤ batas siang) |

Karena problem Anda sangat struktural (blok, split, batas siang, MGMP, guru conflict), maka:

> ❌ Mutation tidak boleh random swap biasa
> ❌ Crossover tidak boleh potong sembarang slot

Kita buat **Smart Mutation** dan **Structure-Preserving Crossover**.

---

# 🔥 STRATEGI BESAR

Karena jam_per_minggu harus selalu sesuai, maka prinsipnya:

> Distribusi jumlah jam jangan dirusak
> Evolusi hanya mengubah posisi (bukan kuantitas)

---

# =========================================================

# 🧬 SMART MUTATION (Constraint-Aware Mutation)

# =========================================================

Kita buat mutasi 3 tahap seperti yang Anda mau.

---

## ✅ TAHAP 1 — Validasi Jam_per_minggu

Untuk satu kelas:

```python
count_subject = Counter(all_slots)
```

Bandingkan dengan:

```python
jam_dictionary[subject_id]
```

---

## ✅ TAHAP 2 — Repair Jika Kelebihan / Kekurangan

Jika:

```
count > jam_dictionary → over
count < jam_dictionary → under
```

Algoritma:

```
for subject_over:
    cari subject_under
    ganti salah satu slot subject_over menjadi subject_under
```

Ini menjaga total slot tetap 36.

---

## ✅ TAHAP 3 — Smart Reshuffle Dalam 1 Hari

Contoh Anda:

```
Senin: [8,2,2,8,2,1,8,1]
```

Tujuan:

* Subject 2 = 3 jam → jadi blok [2,2,2]
* Subject 8 = 3 jam → blok [8,8,8]
* Subject 1 = 2 jam → blok [1,1]

Algoritma:

### 🔹 Langkah 1

Kelompokkan berdasarkan subject

```
{8:3, 2:3, 1:2}
```

### 🔹 Langkah 2

Sort berdasarkan ukuran block (optional)

### 🔹 Langkah 3

Bangun ulang hari secara terurut blok

```
[1,1,2,2,2,8,8,8]
```

Atau bisa random urutan blok, tapi bloknya tetap utuh.

---

# 🎯 Bentuk Final Smart Mutation

### Jenis Mutasi:

### 1️⃣ Repair Mutation

Memperbaiki over/under

### 2️⃣ Block Reordering Mutation

Mengacak urutan blok dalam hari

### 3️⃣ Block Relocation Mutation

Memindahkan 1 blok ke hari lain (untuk memperbaiki split rule)

Contoh:
Subject 4 jam=4 harus [2,2]
Jika dua blok berada di hari sama → pindahkan salah satu ke hari lain.

---

# =========================================================

# 🔁 CROSSOVER YANG TEPAT

# =========================================================

Karena struktur Anda:

```
individu[class][day][slot]
```

Crossover tidak boleh potong di tengah blok.

Saya sarankan 3 opsi berikut:

---

# ✅ OPTION 1 — CLASS-LEVEL CROSSOVER (AMAN)

Parent A:

* kelas 1–13

Parent B:

* kelas 14–27

Child:

* ambil setengah kelas dari A
* setengah dari B

✔ Tidak merusak struktur
✔ Aman jam_per_minggu
✔ Aman blok

Ini paling stabil.

---

# ✅ OPTION 2 — DAY-LEVEL CROSSOVER (RECOMMENDED)

Untuk 1 kelas:

Child:

```
Ambil Senin–Rabu dari Parent A
Ambil Kamis–Jumat dari Parent B
```

Kemudian lakukan repair kecil.

Keuntungan:
✔ Blok tetap utuh (karena blok dalam hari)
✔ Variasi lebih tinggi

---

# ✅ OPTION 3 — BLOCK-BASED CROSSOVER (PALING CANGGIH)

Langkah:

1. Identifikasi semua blok dalam kelas
2. Parent A kirim 50% blok
3. Parent B kirim 50% blok
4. Susun ulang sesuai aturan

Ini paling pintar,
tapi implementasinya lebih kompleks.

---
