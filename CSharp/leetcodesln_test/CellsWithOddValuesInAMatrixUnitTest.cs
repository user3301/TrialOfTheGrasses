// using leetcodesln;
// using Microsoft.VisualStudio.TestTools.UnitTesting;

// namespace leetcodesln_test
// {
//     [TestClass]
//     public class CellsWithOddValuesInAMatrixUnitTest
//     {
//         public CellsWithOddValuesInAMatrix CellsWithOddValuesInAMatrix { get; set; }

//         [TestInitialize]
//         public void BeforeEach()
//         {
//             CellsWithOddValuesInAMatrix = new CellsWithOddValuesInAMatrix();
//         }

//         [TestMethod, Timeout(20_000)]
//         public void TestMethod1()
//         {
//             int n = 2, m = 3;
//             var indices = new int[2][];
//             indices[0] = new[] { 0, 1 };
//             indices[1] = new[] { 1, 1 };
//             var expected = 6;
//             var actual = CellsWithOddValuesInAMatrix.OddCells(n, m, indices);
//             Assert.AreEqual(expected, actual);
//         }

//     }
// }
