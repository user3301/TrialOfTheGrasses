// using leetcodesln;
// using Microsoft.VisualStudio.TestTools.UnitTesting;

// namespace leetcodesln_test
// {
//     [TestClass]
//     public class NumberOfClosedIslandsUnitTest
//     {
//         public NumberOfClosedIslands NumberOfClosedIslands { get; set; }

//         [TestInitialize]
//         public void BeforeEach()
//         {
//             NumberOfClosedIslands = new NumberOfClosedIslands();
//         }

//         [TestMethod, Timeout(20_000)]
//         public void TestMethod1()
//         {
//             var grid = new int[3][];
//             grid[0] = new int[] { 0, 0, 1, 0, 0 };
//             grid[1] = new int[] { 0, 1, 0, 1, 0 };
//             grid[2] = new int[] { 0, 1, 1, 1, 0 };

//             var expected = 1;
//             var actual = NumberOfClosedIslands.ClosedIsland(grid);
//             Assert.AreEqual(expected, actual);
//         }


//         [TestMethod, Timeout(20_000)]
//         public void TestMethod2()
//         {
//             var grid = new int[5][];
//             grid[0] = new int[] { 1, 1, 1, 1, 1, 1, 1, 0 };
//             grid[1] = new int[] { 1, 0, 0, 0, 0, 1, 1, 0 };
//             grid[2] = new int[] { 1, 0, 1, 0, 1, 1, 1, 0 };
//             grid[3] = new int[] { 1, 0, 0, 0, 0, 1, 0, 1 };
//             grid[4] = new int[] { 1, 1, 1, 1, 1, 1, 1, 0 };

//             var expected = 2;
//             var actual = NumberOfClosedIslands.ClosedIsland(grid);
//             Assert.AreEqual(expected, actual);
//         }

//         [TestMethod, Timeout(20_000)]
//         public void TestMethod3()
//         {
//             var grid = new int[5][];
//             grid[0] = new int[] { 1, 1, 1, 1, 1, 1, 1, 0 };
//             grid[1] = new int[] { 1, 0, 0, 0, 0, 1, 1, 0 };
//             grid[2] = new int[] { 1, 0, 1, 0, 1, 1, 1, 0 };
//             grid[3] = new int[] { 1, 0, 0, 0, 0, 1, 0, 1 };
//             grid[4] = new int[] { 1, 1, 1, 1, 1, 1, 1, 0 };

//             var expected = 2;
//             var actual = NumberOfClosedIslands.ClosedIsland(grid);
//             Assert.AreEqual(expected, actual);
//         }

//         [TestMethod, Timeout(20_000)]
//         public void TestMethod4()
//         {
//             // the solution is incorrect
//             Assert.IsTrue(false);
//         }
//     }
// }
