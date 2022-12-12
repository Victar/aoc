package adventofcode.year2022;

import adventofcode.BaseTest;
import adventofcode.GraphUtil;
import lombok.Data;
import org.junit.Ignore;
import org.junit.Test;

import java.util.*;

public class Day12 extends BaseTest {

	public static final int DAY = 12;

	@Ignore @Test public void runDownloadInput() throws Exception {
		downloadInput(DAY);
	}

	@Test public void runSilver() throws Exception {
		runAny(false);
	}

	@Test public void runGold() throws Exception {
		runAny(true);
	}

	public void runAny(boolean isGold) throws Exception {
		final ArrayList<String> data = readStringFromFile("year2022/day" + DAY + "/input.txt");
		final int SIZE_X = data.get(0).length();
		final int SIZE_Y = data.size();

		int END_X = 0;
		int END_Y = 0;
		final List<Node> sourceCandidates = new ArrayList<>();

		final Node[][] NODES = new Node[SIZE_X][SIZE_Y];
		for (int y = 0; y < SIZE_Y; y++) {
			for (int x = 0; x < SIZE_X; x++) {
				int height = data.get(y).charAt(x) - 'a' + 1;
				if (isGold && height == 1) {
					sourceCandidates.add(new Node(x, y, 1));
				}
				if (height == -13) {
					height = 1;
					sourceCandidates.add(new Node(x, y, height));
				}
				if (height == -27) {
					END_X = x;
					END_Y = y;
					height = 'z' - 'a' + 1;

				}
				final Node currentNode = new Node(x, y, height);
				NODES[x][y] = currentNode;
			}
		}

		//set neighbors
		for (int y = 0; y < SIZE_Y; y++) {
			for (int x = 0; x < SIZE_X; x++) {
				for (int l = x - 1; l <= x + 1; l++) {
					for (int m = y - 1; m <= y + 1; m++) {
						if (l == x || m == y) {
							if (l >= 0 && l < SIZE_X && m >= 0 && m < SIZE_Y && (l != x || m != y)) {
								if (NODES[x][y].height + 2 > NODES[l][m].height) {
									NODES[x][y].addDestination(NODES[l][m]);
								}
							}
						}
					}
				}
			}
		}
		int minPath = Integer.MAX_VALUE;
		for (Node node : sourceCandidates) {
			GraphUtil.calculatePath(NODES[node.getX()][node.getY()], 0);
			minPath = Math.min(NODES[END_X][END_Y].getDistance(), minPath);

		}
		System.out.println(minPath);
	}

	public void printArray(final Node[][] array, final int SIZE_X_G, final int SIZE_Y_G, final List<GraphUtil.Node> path) {
		for (int y = 0; y < SIZE_Y_G; y++) {
			for (int x = 0; x < SIZE_X_G; x++) {
				final boolean inPath = path.contains(new Node(x, y, 0));
				if (inPath) {
					System.out.print("(" + array[x][y].getHeight() + ")");
				} else {
					System.out.print(" " + array[x][y].getHeight() + " ");
				}
			}
			System.out.println();
		}
	}

	@Data public class Node implements GraphUtil.Node {

		int x;
		int y;
		int height;
		Map<GraphUtil.Node, Integer> neighbors = new HashMap<>();
		private List<GraphUtil.Node> minPath = new LinkedList<>();
		private Integer distance = Integer.MAX_VALUE;

		public Node(final int x, final int y, final int height) {
			this.x = x;
			this.y = y;
			this.height = height;
		}

		@Override public int getDistance() {
			return this.distance;
		}

		@Override public void setDistance(final int distance) {
			this.distance = distance;
		}

		public void addDestination(final Node destination) {
			this.neighbors.put(destination, 1);//destination.getRisk());
		}

		@Override public boolean equals(final Object o) {
			if (this == o) return true;
			if (o == null || getClass() != o.getClass()) return false;

			final Node node = (Node) o;

			if (this.x != node.x) return false;
			return this.y == node.y;
		}

		@Override public int hashCode() {
			int result = this.x;
			result = 31 * result + this.y;
			return result;
		}

		@Override public void setMinPath(final List<GraphUtil.Node> minPath) {
			this.minPath = minPath;
		}

		@Override public Map<GraphUtil.Node, Integer> getNeighbors() {
			return this.neighbors;
		}
	}

}
