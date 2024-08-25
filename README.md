# go-sql-export

## Export 1 million records from SQL to CSV or Fixed-Length Format Files with high performance

Exporting large volumes of data efficiently from SQL into a CSV or fixed-length format file requires careful consideration of memory management, I/O performance, and database interactions. This article provides an approach to optimize the process in Golang.

![Export data from SQL to CSV or fixed-length format files](https://cdn-images-1.medium.com/max/800/1*IEMXhQXJ0hWZBPL8q2jMNw.png)
### Test Info
- RAM: 12 GB
- Disk: SSD KINGSTON SA400S37240G ATA Device
- Exec File Size: 10M
- Database: PosgreSQL 16
- Total of rows: 1.018.584 rows
- Total of columns: 76 columns
- Power Usage: Very High
- CPU: 15%

<table><thead><tr>
<td><b>Type</b></td>
<td><b>File Size</b></td>
<td><b>Rows</b></td>
<td><b>RAM</b></td>
<td><b>Disk</b></td>
<td><b>Duration</b></td>
<td><b>Description</b></td>
</tr></thead><tbody>

<tr>
<td>Fix Length</td>
<td>1.15 GB</td>
<td>1,018,584</td>
<td>15 M</td>
<td>10.1 M/s</td>
<td>1 min 45 sec</td>
<td>Full scan the table</td>
</tr>

<tr>
<td>CSV</td>
<td>975 MB</td>
<td>1,018,584</td>
<td>15 M</td>
<td>10.1 M/s</td>
<td>1 min 12 sec</td>
<td>Full scan the table</td>
</tr>

<tr>
<td>Fix Length</td>
<td>1.02 GB</td>
<td>905,408</td>
<td>15 M</td>
<td>10.1 M/s</td>
<td>1 min 33 sec</td>
<td>Filter by index on 1 field</td>
</tr>

<tr>
<td>CSV</td>
<td>863 MB</td>
<td>905,408</td>
<td>15 M</td>
<td>10.1 M/s</td>
<td>1 min 3 sec</td>
<td>Filter by index on 1 field</td>
</tr>

<tr>
<td>Fix Length</td>
<td>890 MB</td>
<td>792,232</td>
<td>14 M</td>
<td>9.9 M/s</td>
<td>1 min 23 sec</td>
<td>Filter by index on 1 field</td>
</tr>

<tr>
<td>CSV</td>
<td>764 MB</td>
<td>792,232</td>
<td>14 M</td>
<td>9.9 M/s</td>
<td>55 sec</td>
<td>Filter by index on 1 field</td>
</tr>

<tr>
<td>Fix Length</td>
<td>254 MB</td>
<td>226,352</td>
<td>14 M</td>
<td>9.9 M/s</td>
<td>24 sec</td>
<td>Filter by index on 1 field</td>
</tr>

<tr>
<td>CSV</td>
<td>220 M</td>
<td>226,352</td>
<td>14 M</td>
<td>9.9 M/s</td>
<td>16 sec</td>
<td>Filter by index on 1 field</td>
</tr>

</tbody></table>

### Batch jobs
Differ from online processing:
- Long time running, often at night, after working hours.
- Non-interactive, often include logic for handling errors
- Large volumes of data

### Common Mistakes
- <b>Inefficient Writing to I/O</b>: Large writing to I/O can slow down performance. Writing each record immediately without buffering is inefficient due to the high overhead of repeated I/O operations.
  - <b>Solution</b>: Use "<b>bufio.Writer</b>" for more efficient writing.
- <b>Loading All Data Into Memory</b>: Fetching all records at once can consume a lot of memory, causing the program to slow down or crash. Use streaming with cursors instead.
  - <b>Solution</b>: Loop on each cursor. On each cursor, use bufio.Writer to write to database
- <b>Inefficient Query</b>: Full scan the table. Do not filter on the index.
  - <b>Solution</b>: If you export the whole table, you can scan the full table. If not, you need to filter on the index.

### Implementation
#### Data Reader for SQL
1. Build Query: For efficient query, you need to filter on the index, avoid to scan the full table. In my sample, I created index on field createdDate. In my 6 use cases, I use 4 use cases to filter on indexing field: createdDate.
2. Scan the GO row into an appropriate GO struct:

   We provide a function to map a row to a GO struct. We use gorm tag, so that this struct can be reused for gorm later, with these benefits:
    - Simplifies the process of converting database rows into Go objects.
    - Reduces repetitive code and potential errors in manual data mapping.
    - Enhances code readability and maintainability.
```go
type User struct {
    Id          string     `gorm:"column:id;primary_key" format:"%011s" length:"11"`
    Username    string     `gorm:"column:username" length:"10"`
    Email       string     `gorm:"column:email" length:"31"`
    Phone       string     `gorm:"column:phone" length:"20"`
    Status      bool       `gorm:"column:status" true:"1" false:"0" format:"%5s" length:"5"`
    CreatedDate *time.Time `gorm:"column:createdDate" length:"10" format:"dateFormat:2006-01-02"`
}
```

#### Transformer
Transform a GO struct to a string (CSV or fixed-length format). We created 2 providers already:
- CSV Transformer: read GO tags to transform CSV line.
- Fixed Length Transformer: read GO tags to transform Fixed Length line.

To improve performance, we cache the struct of CSV or Fixed Length Format.

#### File Writer
- It is a wrapper of "<b>bufio.Writer</b>" to buffer writes to the file. This reduces the number of I/O operations.

### Key Aspects to improve performance:
- <b>Streaming</b>: The code uses db.QueryContext to fetch records in a streaming manner. This prevents loading all records into memory at once.
- <b>Memory Management</b>: Since rows are processed one by one, memory usage remains low, even when handling a large number of records.
- <b>Cache Scanning</b>: to improve performance: based on gorm tag, cache column structure when scanning the GO row into an appropriate GO struct.
- <b>Cache Transforming</b>: to improve performance, cache CSV or fixed-length format structure when transforming a GO struct into CSV format or fixed-length for

## Conclusion
In the sample, I tested with 1 million records, I see Postgres still used less than 14M RAM, and my program used about 15M RAM.

So, for this case, we don't need to use LIMIT/OFFSET , as long as we loop on cursor, at each of cursor, we write to file stream.

In the past, I also test with 4 million records, export 4GB, it still works.

### Other Samples:
- [go-hive-export](https://github.com/project-samples/go-hive-export): export data from hive to fix-length or csv file.
- [go-cassandra-export](https://github.com/project-samples/go-cassandra-export): export data from cassandra to fix-length or csv file.
- [go-mongo-export](https://github.com/project-samples/go-mongo-export): export data from mongo to fix-length or csv file.
- [go-firestore-export](https://github.com/project-samples/go-firestore-export): export data from firestore to fix-length or csv file.
